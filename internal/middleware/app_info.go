package middleware

import "github.com/gin-gonic/gin"

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "blog")
		c.Set("app_version", "1.0.0")
		c.Next()
	}
}
