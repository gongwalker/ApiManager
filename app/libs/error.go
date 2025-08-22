package libs

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// ErrorResponse 定义统一的错误响应结构
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// HandleError 统一处理错误
func HandleError(c *gin.Context, err error, statusCode int, message string) {
	if err != nil {
		log.Printf("Error: %v\nStack: %s", err, debug.Stack())
	}

	if message == "" {
		message = "操作失败"
	}

	c.JSON(statusCode, ErrorResponse{
		Code:    statusCode,
		Message: message,
	})
}

// HandleSuccess 统一处理成功响应
func HandleSuccess(c *gin.Context, data interface{}, message string) {
	if message == "" {
		message = "操作成功"
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  message,
		"data": data,
	})
}
