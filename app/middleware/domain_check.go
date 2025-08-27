package middleware

import (
	"ApiManager/app/global"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// DomainCheck 中间件用于限制只允许特定域名访问API
func DomainCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求的Host
		host := c.Request.Host

		// 检查是否是允许的域名
		allowed := false
		for _, allowedHost := range global.AllowedHosts {
			if strings.EqualFold(host, allowedHost) {
				allowed = true
				break
			}
		}

		// 如果不是允许的域名，则拒绝访问
		if !allowed {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		// 继续处理请求
		c.Next()
	}
}
