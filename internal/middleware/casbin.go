package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/limitcool/blog/common"
	"github.com/limitcool/blog/common/errcode"
	response2 "github.com/limitcool/blog/common/response"
	"github.com/limitcool/blog/internal/util"
)

func CheckCasbinAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := response2.NewResponse(c)
		//requestUrl := c.Request.URL.Path
		//method := c.Request.Method
		token := c.GetHeader("token")
		if token == "" {
			response.ToErrorResponse(errcode.InvalidParams)
			c.Abort()
			return
		}
		claims, err := common.ParseToken(token)
		if err != nil {
			response.ToResponse("解码失败,请检查token!")
			c.Abort()
			return
		}
		if claims == nil {
			c.Abort()
			response.ToErrorResponse(errcode.NotFoundToken)
			return
		}

		if claims.Username != util.Md5("admin") {
			fmt.Println(claims.Username)
			response.ToErrorResponse(errcode.NoPermissions)
			c.Abort()
		}
		c.Next()
	}
}
