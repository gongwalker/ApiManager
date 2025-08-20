package routers

import (
	"ApiManager/app/controllers"
	"ApiManager/app/global"
	"ApiManager/app/libs"
	"ApiManager/app/middleware"
	"ApiManager/app/models"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

const sessionSecret = "api_manager_secret"

// 是否登录判定中间件
func authMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("uid")
		role := session.Get("role")
		if uid == nil {
			controllers.JumpLogin(c)
			return
		} else {
			sessionUid, _ := uid.(int)
			sessionRole := role.(int)

			user := models.User{Id: sessionUid}
			curUserInfo, _ := user.GetUserInfoByUid([]string{"id", "role", "isdel"})

			// 签定当前账号是否已被禁用
			if curUserInfo["isdel"].(int64) == 1 {
				controllers.JumpLogin(c)
				return
			}

			// 签定session存放角色值与当前角色值是否一致，若不一致重新登录
			if int(curUserInfo["role"].(int64)) != sessionRole {
				controllers.JumpLogin(c)
				return
			}

			c.Set("UID", uid)
			c.Set("ROLE", role)
		}
		c.Next()
	}
}

// 角色判定中间件 传参role为要求角色
func roleMiddleWare(role int) gin.HandlerFunc {
	return func(c *gin.Context) {
		curRole := libs.GetRole(c) // 当前角色
		switch {
		case curRole == 1: // 若当前为超级管理员
			c.Next()
			return
		case curRole == 2: // 若当前为普通管理员,要求的角色为1的话
			if role == 1 {
				libs.HandleError(c, nil, http.StatusForbidden, "权限不足，无操作权限")
				c.Abort()
				return
			}
		case curRole == 3: //若当前为游客,要求的角色为不为1(超级管理员)或不为2(管理员)
			if role == 1 || role == 2 {
				libs.HandleError(c, nil, http.StatusForbidden, "权限不足，无操作权限")
				c.Abort()
				return
			}
		default: // 未知权限截断
			{
				libs.HandleError(c, nil, http.StatusForbidden, "权限不足，无操作权限")
				c.Abort()
				return
			}
		}
	}
}

func InitRouter() *gin.Engine {
	// 运行模式 release,debug,test
	gin.SetMode(global.GinRunMode)

	// 是否记录运行日志
	if !global.GinWriteLog {
		gin.DefaultWriter = ioutil.Discard
	}

	// 使用自定义中间件
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	// 添加全局API速率限制：每分钟60个请求
	r.Use(middleware.RateLimit(60, time.Minute))
	r.Delims("[[", "]]")

	// session中间件
	r.Use(sessions.Sessions("apiManagerSessionId", getSessionStore()))

	// 指定静态文件
	r.StaticFS("/static", http.Dir("app/static"))
	// 指定模板文件
	r.LoadHTMLGlob("app/views/**/*")
	// 首页
	r.GET("/", authMiddleWare(), controllers.Home)
	// 模块
	r.GET("/m/:module", authMiddleWare(), controllers.Home)

	// api 路由组
	api := r.Group("/api")
	{
		api.Use(authMiddleWare())
		api.GET("/info/:id", controllers.InfoApi)

		// 只有 管理员角色 or 超级管理角色  拥有对接口有增删改权限
		api.POST("/create", roleMiddleWare(2), controllers.DoAddApi)
		api.POST("/edit", roleMiddleWare(2), controllers.DoEditApi)
		api.DELETE("/del/:id", roleMiddleWare(2), controllers.DoDelApi)
		api.POST("/duplicate/:id/:name", roleMiddleWare(2), controllers.DoDuplicateApi)
		api.POST("/sort", roleMiddleWare(2), controllers.SortApi)
	}

	// cate 路由组
	cate := r.Group("/cate")
	{
		cate.Use(authMiddleWare())
		cate.GET("/list", controllers.ListCate)
		cate.GET("/info/:aid", controllers.InfoCate)
		cate.GET("/api/:aid", controllers.ListApi)

		// 只有 超级管理员角色 拥有对接口分类增删改权限
		cate.POST("/add", roleMiddleWare(1), controllers.DoAddCate)
		cate.POST("/edit", roleMiddleWare(1), controllers.DoEditCate)
		cate.DELETE("/del/:aid", roleMiddleWare(1), controllers.DoDelCate)
	}

	// 登录路由组
	login := r.Group("/login")
	{
		login.GET("", controllers.Login)
		login.POST("", controllers.DoLogin)
		login.POST("/exit", controllers.DoExit)
	}

	// 用户管理
	user := r.Group("/user")
	{
		user.Use(authMiddleWare())
		// 用户列表
		user.GET("/list", roleMiddleWare(1), controllers.ListUser)
		// 添加用户
		user.POST("/add", roleMiddleWare(1), controllers.DoAddUser)
		// 编辑用户
		user.POST("/edit", roleMiddleWare(1), controllers.DoEditUser)
		// 重置密码
		user.POST("/resetpwd", controllers.DoResetPwd)
		// 禁用 or 启用用户
		user.POST("/changeStatus", roleMiddleWare(1), controllers.DoChangeStatus)
	}

	// 404 页面设置
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
	return r
}

// 根据配置文件选择session驱动
func getSessionStore() (store sessions.Store) {
	if global.SessionDrive == "redis" {
		return enableSessionRedis()
	}
	return enableSessionCookie()
}

// session redis驱动
func enableSessionRedis() (store sessions.Store) {
	store, err := redis.NewStore(
		global.SessionDriveRedisConfig["size"].(int),
		global.SessionDriveRedisConfig["network"].(string),
		global.SessionDriveRedisConfig["address"].(string),
		global.SessionDriveRedisConfig["password"].(string),
		[]byte(sessionSecret))
	if err != nil {
		log.Fatal("[Redis]", err)
	}
	setOption(store)
	return store
}

// cookie redis驱动
func enableSessionCookie() (store sessions.Store) {
	store = cookie.NewStore([]byte(sessionSecret))
	setOption(store)
	return store
}

// 配置session
func setOption(store sessions.Store) {
	store.Options(sessions.Options{
		MaxAge:   global.SessionOption["max_age"].(int),
		Path:     global.SessionOption["path"].(string),
		Secure:   global.SessionOption["secure"].(bool),
		HttpOnly: global.SessionOption["http_only"].(bool),
		Domain:   global.SessionOption["domain"].(string),
	})
}
