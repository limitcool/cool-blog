package captcha

import (
	"github.com/gin-gonic/gin"
	response2 "github.com/limitcool/blog/common/response"
	"github.com/mojocn/base64Captcha"
	"log"
	"net/http"
)

type Captcha struct {
	ID            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}
type RequestCaptcha struct {
	Id          string `json:"id"`
	CaptchaType string `json:"captcha_type" `
	VerifyValue string `json:"verify_value" `
}

var store = base64Captcha.DefaultMemStore

// GenerateCaptcha 生成验证码
func GenerateCaptcha(c *gin.Context) {
	param := RequestCaptcha{}
	if err := c.ShouldBindJSON(&param); err != nil {
		response2.NewResponse(c).ToResponse(err)
		c.Abort()
		return
	}
	var driver base64Captcha.Driver
	switch param.CaptchaType {
	case "audio":
		driver = base64Captcha.NewDriverAudio(4, "zh")
	case "string":
		//driver = base64Captcha.NewDriverString
		driver = base64Captcha.NewDriverAudio(6, "zh")
	case "math":
		//driver = base64Captcha.NewDriverMath()
		driver = base64Captcha.NewDriverAudio(6, "zh")
	case "chinese":
		//driver = base64Captcha.NewDriverString()
		driver = base64Captcha.NewDriverAudio(6, "zh")
	default:
		driver = base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	}
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		response2.NewResponse(c).ToResponse(err)
		c.Abort()
	}
	param.Id = id
	c.JSON(http.StatusOK, gin.H{
		"captcha_id": id,
		"data":       b64s,
	})
	log.Println(b64s)
}

//
func CaptchaVerify(c *gin.Context) {
	param := RequestCaptcha{}
	// 参数验证
	if err := c.ShouldBindJSON(&param); err != nil {
		response2.NewResponse(c).ToResponse(err)
	}
	if param.VerifyValue == "" {
		response2.NewResponse(c).ToResponse("验证码输入为空")
		c.Abort()
		return
	}
	if store.Verify(param.Id, param.VerifyValue, true) {
		response2.NewResponse(c).ToResponse("验证码输入正确")
		c.Abort()
		return
	}
	log.Printf("id:%s,value:%s", param.Id, param.VerifyValue)
	response2.NewResponse(c).ToResponse("验证码输入错误")
	c.Abort()

}
