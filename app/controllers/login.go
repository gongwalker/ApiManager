package controllers

import (
	"ApiManager/app/Validators"
	"ApiManager/app/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=12"`
}

// 登录界面
func Login(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("uid")

	if uid != nil {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusUnauthorized, "<script>location.href='/'</script>")
		c.Abort()
		return
	}

	c.HTML(http.StatusOK, "login.html", nil)
}

// 登录操作
func DoLogin(c *gin.Context) {
	var loginRequest LoginRequest
	err := c.ShouldBind(&loginRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": Validators.LogInGetError(err.(validator.ValidationErrors)),
		})
		return
	}

	user, err := models.Login(loginRequest.Username, loginRequest.Password)

	if err != nil || user.Id == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Wrong account or password"})
		return
	}

	// 登录成功session设置
	session := sessions.Default(c)
	session.Set("uid", user.Id)
	session.Set("role", user.Role)
	err = session.Save()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Login fail"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Login Success", "data": user})
	return
}

// 退出操作
func DoExit(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	// 注 把maxAge置为负数 清除 redis存放没有意义session-key(相当于把session.clear()过的session_key 设置了expire失效时间)
	// 且 清除客户端cookie
	session.Options(sessions.Options{MaxAge: -1, Domain: "", Secure: false, HttpOnly: true, Path: "/"})
	session.Save()
	// 在session
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
	return
}
