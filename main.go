package main

import (
	"gin-vben-admin/common"
	"gin-vben-admin/initialize"
)

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
	_ = routers.Run(":"+common.CONFIG.System.Port)
}


