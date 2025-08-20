package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimiter 速率限制器
type RateLimiter struct {
	ips   map[string][]time.Time
	mu    sync.Mutex
	limit int           // 限制请求数
	ttl   time.Duration // 时间窗口
}

// NewRateLimiter 创建一个新的速率限制器
func NewRateLimiter(limit int, ttl time.Duration) *RateLimiter {
	return &RateLimiter{
		ips:   make(map[string][]time.Time),
		limit: limit,
		ttl:   ttl,
	}
}

// RateLimit 速率限制中间件
func RateLimit(limit int, ttl time.Duration) gin.HandlerFunc {
	limiter := NewRateLimiter(limit, ttl)

	return func(c *gin.Context) {
		ip := c.ClientIP()

		if !limiter.Allow(ip) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"code": http.StatusTooManyRequests,
				"msg":  "请求过于频繁，请稍后再试",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Allow 检查IP是否允许请求
func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()

	// 清理过期的请求记录
	if _, exists := rl.ips[ip]; exists {
		var validTimes []time.Time
		for _, t := range rl.ips[ip] {
			if now.Sub(t) <= rl.ttl {
				validTimes = append(validTimes, t)
			}
		}
		rl.ips[ip] = validTimes
	}

	// 检查是否超过限制
	if len(rl.ips[ip]) >= rl.limit {
		return false
	}

	// 记录新的请求
	rl.ips[ip] = append(rl.ips[ip], now)

	return true
}
