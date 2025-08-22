package middleware

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

// 日志记录中间件
func Logger() gin.HandlerFunc {
	// 确保日志目录存在
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, 0755)
		if err != nil {
			log.Fatalf("创建日志目录失败: %v", err)
		}
	}

	// 创建日志文件
	logFileName := path.Join(logDir, time.Now().Format("2006-01-02")+".log")
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("打开日志文件失败: %v", err)
	}

	// 设置日志输出
	logger := log.New(logFile, "", log.LstdFlags)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 获取请求体
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 创建自定义响应写入器
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// 记录日志
		responseBody := blw.body.String()
		if len(responseBody) > 1024 {
			responseBody = responseBody[:1024]
		}
		logger.Printf("[%s] %s %s %d %s %s %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			c.Writer.Status(),
			string(requestBody),
			responseBody,
			latency,
		)
	}
}

// 自定义响应写入器，用于捕获响应体
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 重写Write方法
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 重写WriteString方法
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
