package controllers

import (
	"ApiManager/app/Validators"
	"ApiManager/app/libs"
	"ApiManager/app/models"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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

// 上传图片
func UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		libs.HandleError(c, err, http.StatusBadRequest, "Image upload failed")
		return
	}

	// Create a random filename (preserve extension)
	ext := filepath.Ext(file.Filename)
	extNoDot := strings.TrimPrefix(ext, ".")
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		libs.HandleError(c, err, http.StatusInternalServerError, "Failed to generate filename")
		return
	}
	randomStr := hex.EncodeToString(b)
	filename := randomStr + ext

	// Read original file content for DB storage
	src, err := file.Open()
	if err != nil {
		libs.HandleError(c, err, http.StatusInternalServerError, "Failed to read uploaded file")
		return
	}
	defer func() { _ = src.Close() }()
	content, err := io.ReadAll(src)
	if err != nil {
		libs.HandleError(c, err, http.StatusInternalServerError, "Failed to read uploaded file content")
		return
	}

	// Insert upload record into DB
	origName := filepath.Base(file.Filename)
	size := int(file.Size)
	uploadRec := models.Upload{
		FileName:       origName,
		FilePath:       time.Now().Format("200601"),
		UploadFileName: filename,
		Size:           size,
		FileExt:        extNoDot,
		FileContent:    content,
		Addtime:        libs.GetTimeStamp(),
	}
	if err := uploadRec.Insert(); err != nil {
		libs.HandleError(c, err, http.StatusInternalServerError, "Failed to save upload record")
		return
	}

	// Return the public URL via DB-serving endpoint (with file_path year-month)
	publicURL := "/uploads/" + uploadRec.FilePath + "/" + filename
	c.JSON(http.StatusOK, gin.H{"filename": publicURL})
}

/**
 * 展示上传文件（按随机文件名，支持年月路径）
 * 路由：GET /uploads/:filename 或 /uploads/:ym/:filename
 */
func ServeUploadByFilename(c *gin.Context) {
	// 兼容三种写法：/uploads/:filename、/uploads/:ym/:filename、/uploads/*path
	name := c.Param("filename")
	ym := c.Param("ym")
	if name == "" {
		raw := c.Param("path")
		if strings.HasPrefix(raw, "/") {
			raw = raw[1:]
		}
		if raw != "" {
			parts := strings.SplitN(raw, "/", 2)
			if len(parts) == 2 {
				ym = parts[0]
				name = parts[1]
			} else if len(parts) == 1 {
				name = parts[0]
			}
		}
	}
	if name == "" {
		libs.HandleError(c, nil, http.StatusBadRequest, "invalid filename")
		return
	}

	// 规范化文件名与缓存目录
	safeName := filepath.Base(name) // 防止目录穿越
	cacheBase := "app/static/uploads"
	cacheDir := cacheBase
	if ym != "" {
		cacheDir = filepath.Join(cacheBase, filepath.Base(ym))
	}
	cacheFile := filepath.Join(cacheDir, safeName)

	// 1) 命中本地缓存则直接返回
	if fi, err := os.Stat(cacheFile); err == nil && !fi.IsDir() {
		c.File(cacheFile)
		return
	}

	// 2) 未命中缓存，从数据库读取
	var m models.Upload
	rec, err := m.GetByUploadFileName(name)
	if err != nil {
		libs.HandleError(c, err, http.StatusNotFound, "file not found")
		return
	}
	content := rec.FileContent
	if len(content) == 0 {
		libs.HandleError(c, nil, http.StatusNotFound, "file content empty")
		return
	}

	// 3) 将内容写入本地缓存（优先使用 DB 记录的 file_path，其次使用 URL 中的 ym，最后默认目录）
	targetDir := cacheDir
	if rec.FilePath != "" {
		targetDir = filepath.Join(cacheBase, rec.FilePath)
	}
	if err := os.MkdirAll(targetDir, os.ModePerm); err == nil {
		_ = os.WriteFile(filepath.Join(targetDir, safeName), content, 0644)
	}

	// 4) 返回二进制（使用内容嗅探得到 Content-Type）
	ct := http.DetectContentType(content)
	c.Data(http.StatusOK, ct, content)
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
