package controllers

import (
	"ApiManager/app/Validators"
	"ApiManager/app/libs"
	"ApiManager/app/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type AddUserRequest struct {
	LoginName string `json:"login_name" form:"login_name" binding:"required,min=2,max=20"`
	LoginPwd  string `json:"login_pwd" form:"password" binding:"required,min=6,max=16"`
	Role      int    `json:"role" form:"role" binding:"required,min=1,max=3"`
}

type EditUserRequest struct {
	Id        int    `json:"id" form:"id" binding:"required"`
	LoginName string `json:"login_name" form:"login_name" binding:"required,min=2,max=20"`
	Role      int    `json:"role" form:"role" binding:"required,min=1,max=3"`
}

type ReSetPwd struct {
	Id       int    `json:"id" form:"id" binding:"required"`
	LoginPwd string `json:"login_pwd" form:"password" binding:"required,min=6,max=16"`
}

// 用户列表
func ListUser(c *gin.Context) {
	user := models.User{}
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("limit"))
	limit := libs.GetLimitByPage(page, pageSize)

	where := make([]string, 2)
	loginName := libs.Addslashes(c.Query("login_name"))
	role := libs.Addslashes(c.Query("role"))

	if loginName != "" {
		where = append(where, "login_name like '%"+loginName+"%'")
	}
	if role != "" {
		where = append(where, "role='"+role+"'")
	}
	users, total, err := user.Lists(limit, "id desc", where...)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "fail" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": gin.H{"list": users, "count": total}})
}

// 添加用户
func DoAddUser(c *gin.Context) {
	var userRequest AddUserRequest
	err := c.ShouldBind(&userRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": Validators.AddUserGetError(err.(validator.ValidationErrors)),
		})
		return
	}
	user := models.User{}
	user.LoginName = userRequest.LoginName
	user.PassWord = userRequest.LoginPwd
	user.Role = userRequest.Role
	id, err := user.AddUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"msg": "create account fail"})
}

// 编辑用户
func DoEditUser(c *gin.Context) {
	var userRequest EditUserRequest
	err := c.ShouldBind(&userRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": Validators.AddUserGetError(err.(validator.ValidationErrors)),
		})
		return
	}
	user := models.User{}
	user.Id = userRequest.Id
	user.LoginName = userRequest.LoginName
	user.Role = userRequest.Role
	affect, err := user.ModifyUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if affect >= 0 {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"msg": "edit account fail"})
}

// 禁用 or 启用用户
func DoChangeStatus(c *gin.Context) {
	user := models.User{}
	user.Id, _ = strconv.Atoi(c.PostForm("id"))
	user.IsDel, _ = strconv.Atoi(c.PostForm("isdel"))
	res, _ := user.SwitchUser()
	if res > 0 {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"msg": "fail"})
}

// 重置密码
func DoResetPwd(c *gin.Context) {

	var reSetPwd ReSetPwd
	err := c.ShouldBind(&reSetPwd)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": Validators.AddUserGetError(err.(validator.ValidationErrors)),
		})
		return
	}

	if libs.GetRole(c) != 1 && libs.GetUid(c) != reSetPwd.Id {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Insufficient authority, no operation permission"})
		return
	}

	user := models.User{}
	user.Id = reSetPwd.Id
	user.PassWord = reSetPwd.LoginPwd
	affect, err := user.RestUserPwd()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if affect >= 0 {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"msg": "reset password fail"})
}
