package routers

import (
	"ApiManager/app/controllers"
	"ApiManager/app/global"
	"ApiManager/app/libs"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

const sessionSecret = "api_manager_secret"

// 是否登录判定中间件
func authMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("uid")
		role := session.Get("role")
		if uid == nil {
			// 在登录失效的前提下,根据请求的类型(是否为ajax请求) 判断返回返回待登录标识,还是直接跳转到登录页
			if c.GetHeader("X-Requested-With") == "XMLHttpRequest" {
				c.JSON(http.StatusUnauthorized, gin.H{"msg": "Please login"})
			} else {
				c.Header("Content-Type", "text/html; charset=utf-8")
				c.String(http.StatusUnauthorized, "<script>location.href='/login'</script>")
			}
			c.Abort()
			return
		} else {
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
				c.JSON(http.StatusForbidden, gin.H{"msg": "Insufficient authority, no operation permission"})
				c.Abort()
				return
			}
		case curRole == 3: //若当前为游客,要求的角色为不为1(超级管理员)或不为2(管理员)
			if role == 1 || role == 2 {
				c.JSON(http.StatusForbidden, gin.H{"msg": "Insufficient authority, no operation permission"})
				c.Abort()
				return
			}
		default: // 未知权限截断
			{
				c.JSON(http.StatusForbidden, gin.H{"msg": "Insufficient authority, no operation permission"})
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

	r := gin.Default()
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
