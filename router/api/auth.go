package api

import (
	"github.com/gin-gonic/gin"
	"github.com/limitcool/blog/common"
	"github.com/limitcool/blog/internal/service"
	"net/http"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	c.ShouldBind(&param)
	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	token, err := common.GenerateToken(param.Username, param.Password)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	c.JSON(http.StatusOK, token)
}
