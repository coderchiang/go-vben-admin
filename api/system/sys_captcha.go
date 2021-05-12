package system

import (
	"gin-vben-admin/common"
	"gin-vben-admin/common/utils"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Captcha(c *gin.Context) {
	captchaId := captcha.NewLen(common.CONFIG.Captcha.KeyLong)
	c.JSON(http.StatusOK, gin.H{"msg": "验证码获取成功", "data": map[string]interface{}{
			"CaptchaId": captchaId,
			"PicPath":   "/base/captcha/" + captchaId + ".png",
	},
	})
}

func CaptchaImg(c *gin.Context) {
	utils.GinCaptchaServeHTTP(c.Writer, c.Request)
}

