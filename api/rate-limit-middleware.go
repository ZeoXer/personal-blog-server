package api

import (
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type RateLimitMiddleware struct{}

type IPRateLimiter struct {
	ips   map[string]*rate.Limiter
	mutex sync.Mutex
	rate  rate.Limit
	burst int
}

func (i *IPRateLimiter) GetRateLimiter(ip string) *rate.Limiter {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	limiter, exists := i.ips[ip]
	if !exists {
		limiter = rate.NewLimiter(i.rate, i.burst)
		i.ips[ip] = limiter
	}

	return limiter
}

func (m *RateLimitMiddleware) NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	return &IPRateLimiter{
		ips:   make(map[string]*rate.Limiter),
		rate:  r,
		burst: b,
	}
}

func (m *RateLimitMiddleware) RateLimitMiddleware(limiter *IPRateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		l := limiter.GetRateLimiter(ip)

		if !l.Allow() {
			Utils.CJSON(429, "請求過於頻繁，請稍後再試", nil, 0, c)
			c.Abort()
			return
		}

		c.Next()
	}
}

var RateLimitMiddlewareGroup = new(RateLimitMiddleware)
