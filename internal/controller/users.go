package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/limitcool/blog/common/errcode"
	response2 "github.com/limitcool/blog/common/response"
	"github.com/limitcool/blog/internal/service"
)

type UserController struct {
}

func (user UserController) Login(c *gin.Context) {
	// 获取参数和参数校验
	param := service.LoginRequest{}
	if errs := c.ShouldBindJSON(&param); errs != nil {
		response2.NewResponse(c).ToErrorResponse(errcode.InvalidParams)
	}
	svc := service.New(c.Request.Context())
	err := svc.CheckLogin(&param)
	if err != nil {
		response2.NewResponse(c).ToResponse(err.Error())
		return
	}

	response2.NewResponse(c).ToErrorResponse(errcode.Success)

}

func (user UserController) Register(c *gin.Context) {
	param := service.RegisterRequest{}
	if err := c.ShouldBindJSON(&param); err != nil {
		response2.NewResponse(c).ToErrorResponse(errcode.InvalidParams)
	}
	svc := service.New(c.Request.Context())
	err := svc.CheckRegister(&param)
	if err != nil {
		response2.NewResponse(c).ToResponse(err.Error())
		return
	}
	response2.NewResponse(c).ToErrorResponse(errcode.Success)
}
func NewUserController() UserController {
	return UserController{}
}
