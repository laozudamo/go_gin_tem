package controller

import (
	response "goGinTem/Response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

// base64Captcha  缓存对象
var store = base64Captcha.DefaultMemStore

func GetCaptcha(ctx *gin.Context) {
	//
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	// b64s是图片的base64编码
	id, b64s, err := cp.Generate()
	if err != nil {
		zap.S().Errorf("生成验证码错误,:%s ", err.Error())
		response.Err(ctx, http.StatusInternalServerError, 500, "生成验证码错误", "")
		return
	}
	response.Success(ctx, 200, "生成验证码成功", gin.H{
		"captchaId": id,
		"picPath":   b64s,
	})
}
