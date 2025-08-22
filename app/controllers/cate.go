package controllers

import (
	"ApiManager/app/Validators"
	"ApiManager/app/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type AddCateRequest struct {
	Cname string `json:"cname" form:"cname" binding:"required,max=50"`
	Cdesc string `json:"cdesc" form:"cdesc" binding:"required,max=200"`
}

type EditCateRequest struct {
	AddCateRequest
	Aid   int    `json:"aid" form:"aid" binding:"required"`
	Csort string `json:"csort" form:"csort" binding:"omitempty,number,gte=0,lte=9999"`
}

// 分类列表
func ListCate(c *gin.Context) {
	cate := models.Cate{}
	cates := cate.Lists()
	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": cates})
}

// 添加分类
func DoAddCate(c *gin.Context) {
	var cateRequest AddCateRequest
	err := c.ShouldBind(&cateRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": Validators.AddCateGetError(err.(validator.ValidationErrors)),
		})
		return
	}
	cate := models.Cate{Cname: cateRequest.Cname, Cdesc: cateRequest.Cdesc}
	err = cate.Add()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "create collection failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "create collection success"})
}

// 分类详情(添加界面)
func InfoCate(c *gin.Context) {
	aid, _ := strconv.Atoi(c.Param("aid"))
	cate := models.Cate{Aid: aid}
	info, _ := cate.Info()
	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": info})
}

// 编辑分类
func DoEditCate(c *gin.Context) {
	var cateRequest EditCateRequest
	err := c.ShouldBind(&cateRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": Validators.EditCateGetError(err.(validator.ValidationErrors)),
		})
		return
	}

	ord, _ := strconv.Atoi(cateRequest.Csort)
	cate := models.Cate{Cname: cateRequest.Cname, Cdesc: cateRequest.Cdesc, Aid: cateRequest.Aid, Csort: ord}
	err = cate.Edit()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "edit collection failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "edit collection success"})
}

// 删除分类
func DoDelCate(c *gin.Context) {
	aid, _ := strconv.Atoi(c.Param("aid"))
	cate := models.Cate{Aid: aid}
	err := cate.Delete()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Delete collection failed" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Delete collection success"})
}
