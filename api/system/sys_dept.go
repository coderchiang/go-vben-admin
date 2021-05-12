package system

import (
	"gin-vben-admin/dto"
	"gin-vben-admin/middleware"
	"gin-vben-admin/service"
	"github.com/gin-gonic/gin"
)



func GetDeptTreeAll(c *gin.Context) {
		var q dto.QuerySysDept
		if err := c.ShouldBind(&q); err != nil {
			middleware.ResponseFail(c,201,err.Error())
			return
		}
	menus, _, err := service.GetDeptTree(q)
	if err != nil {
		middleware.ResponseFail(c,202,err.Error())
	} else {
		middleware.ResponseSucc(c,"获取部门成功",menus)
	}
	return
}




func CreateDept(c *gin.Context) {
	var dept dto.SysDept
	err := c.ShouldBindJSON(&dept)
	if err != nil {
		middleware.ResponseFail(c,201,err.Error())
		return
	}
		err = service.SaveDept(&dept)

	if err != nil {
		middleware.ResponseFail(c,202,err.Error())
	} else {
		middleware.ResponseSucc(c,"添加部门成功", true)
		return
	}

}

func UpdateDept(c *gin.Context) {
	var dept dto.SysDept
	err := c.ShouldBindJSON(&dept)
	if err != nil {
		middleware.ResponseFail(c,201,err.Error())
		return
	}
	err = service.SaveDept(&dept)

	if err != nil {
		middleware.ResponseFail(c,202,err.Error())
	} else {
		middleware.ResponseSucc(c,"添加部门成功", true)
		return
	}

}


func DelDept(c *gin.Context) {
	var delParam dto.QuerySysDept
	err := c.ShouldBindJSON(&delParam)
	if err!=nil {

		middleware.ResponseFail(c,201,err.Error())
		return
	}
	 err1:= service.DelDept(delParam.ID)
	if err1 != nil {
		middleware.ResponseFail(c,202,err1.Error())
	} else {
		middleware.ResponseSucc(c,"删除部门成功", true)
	}
	return
}
