package initialize

import (
	"gin-vben-admin/common"
	"gin-vben-admin/middleware"
	"gin-vben-admin/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)


func InitRouters() (Router *gin.Engine) {
	Router = gin.Default()

	Router.Use(middleware.Cors())
	common.LOG.Info("use middleware cors logger recovery")

	Router.Static("/dist", "./web/dist")
	//Router.LoadHTMLGlob("templates/*")
	Router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/dist")
	})

	ApiV1Group := Router.Group("api")

	routers.InitBaseRouter(ApiV1Group)
	routers.InitUserRouter(ApiV1Group)
	routers.InitRoleRouter(ApiV1Group)
	routers.InitMenuRouter(ApiV1Group)
	routers.InitOpLogRouter(ApiV1Group)
	routers.InitDeptRouter(ApiV1Group)
	common.LOG.Info("routers register success")

	return
}