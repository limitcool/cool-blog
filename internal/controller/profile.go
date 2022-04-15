package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/limitcool/blog/internal/service"
)

type ProfileController struct {
}

func (p ProfileController) Create(c *gin.Context) {
	param := service.ProfileRequest{}
	_ = c.ShouldBindJSON(&param)
	svc := service.New(c.Request.Context())
	id, err := svc.ProfileCreate(&param)
	if err != nil {
		c.JSON(404, err)
		return
	}
	c.JSON(200, id)
}

func NewProfileController() ProfileController {
	return ProfileController{}
}
