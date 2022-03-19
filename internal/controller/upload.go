package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/limitcool/blog/common/errcode"
	response2 "github.com/limitcool/blog/common/response"
	"github.com/limitcool/blog/common/upload"
	"github.com/limitcool/blog/internal/service"
	"log"
	"strconv"
)

type Uplaod struct {
}

func NewUpload() Uplaod {
	return Uplaod{}
}
func (u Uplaod) UploadFile(c *gin.Context) {
	response := response2.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	fileType, _ := strconv.Atoi(c.PostForm("type"))

	if fileHeader == nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		log.Printf("svc.UploadFile err:%v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
