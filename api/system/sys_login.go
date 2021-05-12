package system

import (
	"gin-vben-admin/dto"
	"gin-vben-admin/middleware"
	"gin-vben-admin/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Login(c *gin.Context){
	var loginForm dto.UserLogin

	if err := c.ShouldBindJSON(&loginForm); err != nil {
		middleware.ResponseFail(c,201,err.Error())
		service.CreatOpLog(c, loginForm.Username,time.Duration(1),err.Error())
		return
	}
	if true {
	//if captcha.VerifyString(loginForm.CaptchaId, loginForm.Captcha) {
		user, msg, isPass := service.LoginCheck(loginForm.Username,loginForm.Password)
		if !isPass {
			middleware.ResponseFail(c,202,msg)
			service.CreatOpLog(c, loginForm.Username,time.Duration(1),msg)
		}else {
            //用户返回值
			var res dto.LoginSucc
			res.UserInfo.ID =strconv.Itoa(user.ID)
			//res.UserInfo.AvatarUrl=user.AvatarUrl
			res.UserInfo.RoleId =strconv.Itoa(user.RoleId)
			res.UserInfo.Status =user.Status
			res.UserInfo.Username=user.Username
			token, msg, ok := middleware.GenerateToken(&res.UserInfo)
			res.Token =token

			if !ok {
				middleware.ResponseFail(c,203,msg)
				service.CreatOpLog(c, loginForm.Username,time.Duration(1),msg)
			}else {
				middleware.ResponseSucc(c,msg,res)
				service.CreatOpLog(c, loginForm.Username,time.Duration(1),msg)
			}
		}
	}else {
		middleware.ResponseFail(c,204,"验证码错误")
		service.CreatOpLog(c, loginForm.Username,time.Duration(1),"captcha verify error!")
	}
	return
}

func LogOut(c *gin.Context) {
	middleware.ResponseSucc(c,  "退出成功",nil)
	Username:=middleware.GetClaims(c).Username
	service.CreatOpLog(c, Username,time.Duration(1),"logout success!")
}

