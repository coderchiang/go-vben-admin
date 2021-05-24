package system

import (
	"gin-vben-admin/common"
	"gin-vben-admin/common/utils"
	"gin-vben-admin/middleware"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func Captcha(c *gin.Context) {
	captchaId := captcha.NewLen(common.CONFIG.Captcha.KeyLong)
	middleware.ResponseSucc(c,"验证码获取成功",map[string]interface{}{
		"CaptchaId": captchaId,
		"CaptchaSrc":"http://"+c.Request.Host+c.Request.URL.Path +"/"+ captchaId + ".png",
	})
}

func CaptchaImg(c *gin.Context) {
	utils.CaptchaServeHTTP(c.Writer, c.Request)
}

