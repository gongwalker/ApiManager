package controllers

import (
	"ApiManager/app/Validators"
	"ApiManager/app/libs"
	"ApiManager/app/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type apiRequest struct {
	Aid    int    `json:"aid" form:"aid" binding:"required,max=50"`
	Num    string `json:"num" form:"num" binding:"required,max=50"`
	Name   string `json:"name" form:"name" binding:"required,max=200"`
	Url    string `json:"url" form:"url" binding:"required,max=200"`
	Des    string `json:"des" form:"des" binding:"omitempty,max=200"`
	Type   string `json:"type" form:"type" binding:"omitempty"`
	Re     string `json:"re" form:"re" binding:"omitempty,max=100000"`
	St     string `json:"st" form:"st" binding:"omitempty,max=100000"`
	Memo   string `json:"memo" form:"memo" binding:"omitempty,max=100000"`
	Extend string `json:"extend" form:"extend" binding:"omitempty,max=100000"`
	params
}

type params struct {
	ParamName    []string `json:"param_name" form:"p[name][]" binding:"param_name"`
	ParamType    []string `json:"param_type" form:"p[paramType][]" binding:"param_type"`
	ParamCate    []string `json:"param_cate" form:"p[type][]"`
	ParamDefault []string `json:"param_default" form:"p[default][]"`
	ParamDes     []string `json:"param_des" form:"p[des][]"`
}

// 某个分类的下的接口列表
func ListApi(c *gin.Context) {
	aid, _ := strconv.Atoi(c.Param("aid"))
	api := models.Api{Aid: aid}
	cate := models.Cate{Aid: aid}
	cateInfo, _ := cate.Info()
	apis := api.Lists()
	type cateContainer struct {
		Apis *[]models.Api `json:"apis"`
		Info *models.Cate  `json:"info"`
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": cateContainer{Apis: &apis, Info: &cateInfo}})
}

// 接口详情(添加界面)
func InfoApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		libs.HandleError(c, err, http.StatusBadRequest, "无效的接口ID")
		return
	}

	api := models.Api{Id: id}
	info, err := api.Info()
	if err != nil {
		libs.HandleError(c, err, http.StatusInternalServerError, "获取接口信息失败")
		return
	}

	libs.HandleSuccess(c, info, "获取接口信息成功")
}

// 添加接口
func DoAddApi(c *gin.Context) {
	var apiRequest apiRequest
	err := c.ShouldBind(&apiRequest)
	if err != nil {
		libs.HandleError(c, err, http.StatusUnprocessableEntity, Validators.ApiGetError(err.(validator.ValidationErrors)))
		return
	}

	pars := params{
		ParamName:    apiRequest.ParamName,
		ParamType:    apiRequest.ParamType,
		ParamCate:    apiRequest.ParamCate,
		ParamDefault: apiRequest.ParamDefault,
		ParamDes:     apiRequest.ParamDes,
	}
	uid := libs.GetUid(c)
	LastTime := libs.GetTimeStamp()
	jsonBytes, err := json.Marshal(pars)
	if err != nil {
		libs.HandleError(c, err, http.StatusInternalServerError, "参数序列化失败")
		return
	}

	api := models.Api{
		Aid:       apiRequest.Aid,
		Num:       apiRequest.Num,
		Name:      apiRequest.Name,
		Des:       apiRequest.Des,
		Url:       apiRequest.Url,
		Type:      apiRequest.Type,
		Re:        apiRequest.Re,
		St:        apiRequest.St,
		Memo:      apiRequest.Memo,
		Parameter: string(jsonBytes),
		Lastuid:   uid,
		Lasttime:  LastTime,
	}
	err = api.Add()
	if err != nil {
		libs.HandleError(c, err, http.StatusInternalServerError, "创建接口失败")
		return
	}
	libs.HandleSuccess(c, nil, "创建接口成功")
}

// 编辑接口
func DoEditApi(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil || id == 0 {
		libs.HandleError(c, err, http.StatusBadRequest, "无效的接口ID")
		return
	}

	var apiRequest apiRequest
	err = c.ShouldBind(&apiRequest)
	if err != nil {
		libs.HandleError(c, err, http.StatusUnprocessableEntity, Validators.ApiGetError(err.(validator.ValidationErrors)))
		return
	}

	pars := params{
		ParamName:    apiRequest.ParamName,
		ParamType:    apiRequest.ParamType,
		ParamCate:    apiRequest.ParamCate,
		ParamDefault: apiRequest.ParamDefault,
		ParamDes:     apiRequest.ParamDes,
	}
	uid := libs.GetUid(c)
	LastTime := libs.GetTimeStamp()
	jsonBytes, err := json.Marshal(pars)
	if err != nil {
		libs.HandleError(c, err, http.StatusInternalServerError, "参数序列化失败")
		return
	}

	api := models.Api{
		Id:        id,
		Aid:       apiRequest.Aid,
		Num:       apiRequest.Num,
		Name:      apiRequest.Name,
		Des:       apiRequest.Des,
		Url:       apiRequest.Url,
		Type:      apiRequest.Type,
		Re:        apiRequest.Re,
		St:        apiRequest.St,
		Memo:      apiRequest.Memo,
		Parameter: string(jsonBytes),
		Lastuid:   uid,
		Lasttime:  LastTime,
	}
	err = api.Edit()
	if err != nil {
		libs.HandleError(c, err, http.StatusInternalServerError, "编辑接口失败")
		return
	}
	libs.HandleSuccess(c, nil, "编辑接口成功")
}

// 删除接口
func DoDelApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		libs.HandleError(c, err, http.StatusBadRequest, "无效的接口ID")
		return
	}

	api := models.Api{Id: id}
	err = api.Delete()
	if err != nil {
		libs.HandleError(c, err, http.StatusInternalServerError, "删除接口失败")
		return
	}
	libs.HandleSuccess(c, nil, "删除接口成功")
}

// 复制接口
func DoDuplicateApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		libs.HandleError(c, err, http.StatusBadRequest, "无效的接口ID")
		return
	}

	name := c.Param("name")
	if name == "" {
		libs.HandleError(c, nil, http.StatusBadRequest, "接口名称不能为空")
		return
	}

	uid := libs.GetUid(c)
	LastTime := libs.GetTimeStamp()
	api := models.Api{Id: id, Name: name, Lastuid: uid, Lasttime: LastTime}
	err = api.DuplicateApi()
	if err != nil {
		libs.HandleError(c, err, http.StatusInternalServerError, "复制接口失败")
		return
	}
	libs.HandleSuccess(c, nil, "复制接口成功")
}

// 排序
func SortApi(c *gin.Context) {
	ids := c.PostFormArray("ids[]")
	if len(ids) == 0 {
		libs.HandleError(c, nil, http.StatusBadRequest, "排序参数不能为空")
		return
	}

	api := models.Api{}
	err := api.Sort(ids)
	if err != nil {
		libs.HandleError(c, err, http.StatusInternalServerError, "接口排序失败")
		return
	}
	libs.HandleSuccess(c, nil, "接口排序成功")
}
