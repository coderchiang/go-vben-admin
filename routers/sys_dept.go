package routers

import (
	v1 "gin-vben-admin/api/system"
	"gin-vben-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDeptRouter(Router *gin.RouterGroup)  {
	MenuRouter := Router.Group("system").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.Logger())
	{
		MenuRouter.GET("/dept/list", v1.GetDeptTreeAll)
		MenuRouter.DELETE("/dept/del", v1.DelDept)
		MenuRouter.POST("/dept/add", v1.CreateDept)
		MenuRouter.PUT("/dept/edit", v1.UpdateDept)
	}
}
