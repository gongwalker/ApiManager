package middleware

import (
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 检查IP是否为内网IP
func isPrivateIP(ip net.IP) bool {
	// 检查是否为本地回环地址 (127.0.0.0/8)
	if ip.IsLoopback() {
		return true
	}

	// 检查是否为内网IP地址
	// 10.0.0.0/8
	if ip[0] == 10 {
		return true
	}
	// 172.16.0.0/12
	if ip[0] == 172 && ip[1] >= 16 && ip[1] <= 31 {
		return true
	}
	// 192.168.0.0/16
	if ip[0] == 192 && ip[1] == 168 {
		return true
	}

	// 特定的内网IP段 (如果有特殊需求)
	// 例如: 10.0.7.0/24
	if ip[0] == 10 && ip[1] == 0 && ip[2] == 7 {
		return true
	}

	return false
}

// 获取客户端真实IP
func getRealIP(c *gin.Context) string {
	// 尝试从X-Forwarded-For头获取
	xForwardedFor := c.Request.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		// X-Forwarded-For可能包含多个IP，第一个是客户端真实IP
		ips := strings.Split(xForwardedFor, ",")
		if len(ips) > 0 {
			clientIP := strings.TrimSpace(ips[0])
			if clientIP != "" {
				return clientIP
			}
		}
	}

	// 尝试从X-Real-IP头获取
	xRealIP := c.Request.Header.Get("X-Real-IP")
	if xRealIP != "" {
		return xRealIP
	}

	// 直接从请求中获取
	return c.ClientIP()
}

// DomainCheck 中间件用于限制只允许内网IP访问API
func DomainCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端IP
		clientIP := getRealIP(c)

		// 解析IP地址
		ip := net.ParseIP(clientIP)
		if ip == nil {
			// 无法解析IP地址，拒绝访问
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		// 检查是否为内网IP
		if !isPrivateIP(ip) {
			// 不是内网IP，拒绝访问
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		// 是内网IP，允许访问
		c.Next()
	}
}
