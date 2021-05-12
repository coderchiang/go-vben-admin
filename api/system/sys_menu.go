package system

import (
	"gin-vben-admin/dto"
	"gin-vben-admin/middleware"
	"gin-vben-admin/service"
	"github.com/gin-gonic/gin"
	//"strconv"
)



func GetMenuTreeAll(c *gin.Context) {
		var q dto.QuerySysMenu
		if err := c.ShouldBind(&q); err != nil {
			middleware.ResponseFail(c,201,err.Error())
			return
		}
	menus, _, err := service.GetMenuTree(q)
	if err != nil {
		middleware.ResponseFail(c,202,err.Error())
	} else {
		middleware.ResponseSucc(c,"获取全部菜单成功",menus)
	}
	return
}

func CreatMenu(c *gin.Context) {
	var menu dto.SysMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		middleware.ResponseFail(c,201,err.Error())
		return
	}
	err = service.SaveMenu(&menu)
	if err != nil {
		middleware.ResponseFail(c,202,err.Error())
	} else {
		middleware.ResponseSucc(c,"添加菜单成功", true)
		return
	}

}



func UpdateMenu(c *gin.Context) {
	var menu dto.SysMenu


	err := c.ShouldBindJSON(&menu)
	if err != nil {
		middleware.ResponseFail(c,201,err.Error())
		return
	}
		err = service.SaveMenu(&menu)
	if err != nil {
		middleware.ResponseFail(c,202,err.Error())
	} else {
		middleware.ResponseSucc(c,"更新菜单成功", true)
		return
	}

}



func DeleteMenu(c *gin.Context) {
	var delParam dto.QuerySysMenu
	err := c.ShouldBindJSON(&delParam)
	if err!=nil {

		middleware.ResponseFail(c,201,err.Error())
		return
	}
	 err1:= service.DelMenu(delParam.ID)
	if err1 != nil {
		middleware.ResponseFail(c,202,err1.Error())
	} else {
		middleware.ResponseSucc(c,"删除菜单成功", true)
	}
	return
}
