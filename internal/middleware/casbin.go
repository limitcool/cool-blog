package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limitcool/blog/common"
	"github.com/limitcool/blog/common/errcode"
	response2 "github.com/limitcool/blog/common/response"
	"github.com/limitcool/blog/global"
	"github.com/limitcool/blog/internal/dao"
	"log"
)

func CheckCasbinAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := response2.NewResponse(c)
		requestUrl := c.Request.URL.Path
		log.Println(requestUrl)
		method := c.Request.Method
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

			response.ToErrorResponse(errcode.NotFoundToken)
			c.Abort()
			return
		}
		// 判断用户是否是admin
		//if claims.Username != "admin" {
		//	response.ToErrorResponse(errcode.NoPermissions)
		//	c.Abort()
		//}
		var GDao dao.Dao
		roleid := GDao.GetRoleId(claims.Username)
		log.Println("roleid:", roleid)
		ok, err := global.Enforcer.Enforce(roleid, requestUrl, method)
		if err != nil {
			c.Abort()
			response.ToResponse(err)
			return
		}
		if !ok {
			response.ToErrorResponse(errcode.NoPermissions)
			c.Abort()
		}

		c.Next()
	}
}
