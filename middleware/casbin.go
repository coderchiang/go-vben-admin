package middleware

import (
	"gin-vben-admin/common"
	"github.com/gin-gonic/gin"
	"strings"
)

//拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user:=GetClaims(c)
		//获取用户的角色
		sub := user.RoleId
		//获取请求的URI
		obj := strings.Replace(c.Request.URL.Path, "/api", "", 1)
		//获取请求方法
		act := c.Request.Method

		_=common.CASBIN.LoadPolicy()
			ok,err:= common.CASBIN.EnforceSafe(sub, obj, act)
			if common.CONFIG.System.Env == "develop" ||ok{

					c.Next()

		} else {
			if err!=nil{
				common.LOG.Warn("Casbin规则错误"+"\t"+ err.Error())
			}
				common.LOG.Warn(user.Username + "\t" + obj + "\t" + act )
			ResponseFail(c, 403, "非法请求\t权限不足")
			c.Abort()
		}
	}
}
