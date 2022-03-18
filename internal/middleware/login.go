package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 必须登录判断 中间件
func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// _为返回值,status为返回状态(bool类型)
		if _, status := c.GetQuery("token"); !status {
			c.String(http.StatusUnauthorized, "缺少token参数!")
			// c.Abort() 停止响应
			c.Abort()
		} else {
			// c.Next() 继续下一步
			c.Next()
		}
	}
}

func Cookie() {

}

// 定义一个JWTAuth的中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过http header中的token解析来认证
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "请求未携带token,无权限访问",
				"data":   nil,
			})
			c.Abort()
			return
		}
		log.Println("get token: ", token)
		log.Println("token is ok")
		c.Next()
	}
}
