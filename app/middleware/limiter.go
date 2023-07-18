package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
	"time"
)

// @Title  middleware
// @Description  MyGO
// @Author  WeiDa  2023/7/18 11:59
// @Update  WeiDa  2023/7/18 11:59

type RateLimiter struct {
	bucket map[string]*TokenBucket
	mutex  sync.Mutex
}

type TokenBucket struct {
	rate       float64   // 速率，单位：令牌/秒
	capacity   float64   // 令牌桶容量
	tokens     float64   // 当前令牌数量
	lastUpdate time.Time // 上次更新时间
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		bucket: make(map[string]*TokenBucket),
	}
}

// LimitHandler 接口IP限流，黑名单&白名单的实例
func LimitHandler(maxConn, rate float64) gin.HandlerFunc {
	limiter := NewRateLimiter()
	return func(c *gin.Context) {
		if limiter.allowIP(c, maxConn, rate) {
			c.Next()
		} else {
			c.JSON(http.StatusTooManyRequests, gin.H{"message": "请求过于频繁,请稍后再试!!!"})
			c.Abort()
			return
		}
	}
}

// allowIP 检查IP是否允许访问 接口IP限流，黑名单&白名单的实例
func (rl *RateLimiter) allowIP(c *gin.Context, maxConn, rate float64) bool {

	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	ip := getRealIp(c)
	bucket, exists := rl.bucket[ip]
	if !exists {
		// 初始化令牌桶
		bucket = &TokenBucket{
			rate:       rate,    // 每秒生成10个令牌
			capacity:   maxConn, // 令牌桶容量为10个
			tokens:     0,       // 初始时令牌桶为0
			lastUpdate: time.Now(),
		}
		rl.bucket[ip] = bucket
	}

	// 计算时间间隔，并根据速率生成令牌
	now := time.Now()
	elapsed := now.Sub(bucket.lastUpdate).Seconds()
	tokensToAdd := elapsed * bucket.rate

	// 更新令牌桶状态
	if tokensToAdd > 0 {
		bucket.tokens = bucket.tokens + tokensToAdd
		if bucket.tokens > bucket.capacity {
			bucket.tokens = bucket.capacity
		}
		bucket.lastUpdate = now
	}

	// 检查令牌数量是否足够
	if bucket.tokens >= 1 {
		bucket.tokens--
		return true
	}
	return false
}

// getRealIp 得到请求的真实IP
func getRealIp(c *gin.Context) (ip string) {
	ip = c.Request.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = c.Request.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = c.Request.RemoteAddr
	}
	log.Printf("Request from IP %s \n", ip)
	return
}
