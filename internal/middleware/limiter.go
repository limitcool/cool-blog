package middleware

import (
	"github.com/gin-gonic/gin"
	limiter "github.com/limitcool/blog/common/limiter"
	"log"
	"net/http"
)

func MaxAllowed(limitValue int64) func(c *gin.Context) {
	limiter := limiter.NewLimiter(limitValue)
	log.Println("limiter.SetMax:", limitValue)
	// 返回限流逻辑
	return func(c *gin.Context) {
		if !limiter.Ok() {
			c.AbortWithStatus(http.StatusServiceUnavailable) //超过每秒200，就返回503错误码
			return
		}
		c.Next()
	}
}
