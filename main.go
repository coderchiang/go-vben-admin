package main

import (
	"gin-vben-admin/common"
	_ "gin-vben-admin/docs"
	"gin-vben-admin/initialize"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
// @title go vben admin API
// @version 1.0
// @description  Golang api of admin
// @termsOfService https://github.com/coderchiang/go-vben-admin

// @contact.name chrischiang
// @contact.url http://2wm.top
// @contact.email 2501170033@qq.com

//@host 127.0.0.1:80
func main()  {
	//初始华配置
	initialize.InitConf()
	// 初始化日志
	initialize.InitLog()
	//初始化redis
	initialize.InitCache()
	defer common.CACHE.Close()
	//初始化数据库
	initialize.InitDb()
	defer common.DB.Close()
	// 初始化路由
	initialize.InitCasbin()
	routers := initialize.InitRouters()
	routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	_ = routers.Run(":"+common.CONFIG.System.Port)
}


