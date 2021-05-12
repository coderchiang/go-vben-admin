package routers

import (
	"github.com/gin-gonic/gin"
	v1 "gin-vben-admin/api/system"
	"gin-vben-admin/middleware"
)

func InitOpLogRouter(Router *gin.RouterGroup)  {
	OpLogRouter := Router.Group("system").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		OpLogRouter.GET("/log/list", v1.GetOperLogList)
		//OpLogRouter.GET("/getOperLogListByIds", v1.GetOperLogListByIds)
		OpLogRouter.DELETE("/log/del", v1.DeleteLog)
		OpLogRouter.DELETE("/log/del_batch", v1.BatchDeleteLog)
		//OpLogRouter.DELETE("/oplog/clean", v1.CleanOperLog)
	}
}

