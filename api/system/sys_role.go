package system

import (
	"gin-vben-admin/dto"
	"gin-vben-admin/middleware"
	"gin-vben-admin/service"
	"github.com/gin-gonic/gin"
)


func GetRoleList(c *gin.Context)  {
	var q dto.QuerySysRole
	if err := c.ShouldBind(&q); err != nil{
		middleware.ResponseFail(c,201,err.Error())
		return
	}
	roles,total, err := service.GetRoleList(q)
	if err!=nil {
		middleware.ResponseFail(c,202,err.Error())
		return
	}else {
		if q.Page==""||q.PageSize==""{
			middleware.ResponseSucc(c, "获取角色列表",roles)
		}else{
			middleware.ResponseSucc(c, "获取角色列表",map[string]interface{}{
				"items":roles,
				"total":total,
			})
		}

	}
	return
}

func CreateRole(c *gin.Context)  {
	var role dto.SysRole
	if err := c.ShouldBindJSON(&role); err != nil{
		middleware.ResponseFail(c,201,err.Error())
		return
	}

	err := service.SaveRole(role)
	if err!=nil {
		middleware.ResponseFail(c,203,err.Error())
		return
	}else {
		middleware.ResponseSucc(c, "添加角色成功",true)
	}


	return
}
func UpdateRole(c *gin.Context)  {
	var role dto.SysRole
	if err := c.ShouldBindJSON(&role); err != nil{
		middleware.ResponseFail(c,201,err.Error())
		return
	}

		err := service.SaveRole(role)
		if err!=nil {
			middleware.ResponseFail(c,203,err.Error())
			return
		}else {
			middleware.ResponseSucc(c, "修改角色成功",true)
		}


	return
}


func DelRole(c *gin.Context)  {
	var Param dto.QuerySysRole
	err := c.ShouldBindJSON(&Param)
	if err!=nil {

		middleware.ResponseFail(c,201,err.Error())
		return
	}
	err1:= service.DelRole(Param.ID)
	if err1 != nil {
		middleware.ResponseFail(c,202,err1.Error())
	} else {
		middleware.ResponseSucc(c,"删除角色成功", true)
	}
	return
}

