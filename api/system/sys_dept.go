package system

import (
	"gin-vben-admin/dto"
	"gin-vben-admin/middleware"
	"gin-vben-admin/service"
	"github.com/gin-gonic/gin"
)


// @Summary 部门列表
// @Description 获取部门列表
// @Tags 部门管理
// @accept json
// @Produce  json
// @Param Authorization header   string true "token"
// @Param id query int false "id"
// @Param deptName query string false "部门名"
// @Param status query string false "状态"
// @Param page query string false "当前页"
// @Param pageSize query string false "每页条数"
// @Success 200 {object} middleware.Response{result=dto.SysDeptOutput{items=[]dto.SysDept}}
// @Router /api/system/dept/list [get]
func GetDeptTreeAll(c *gin.Context) {
		var q dto.QuerySysDept
		if err := c.ShouldBind(&q); err != nil {
			middleware.ResponseFail(c,201,err.Error())
			return
		}
	menus, err := service.GetDeptTree(q)
	if err != nil {
		middleware.ResponseFail(c,201,err.Error())
	} else {
		middleware.ResponseSucc(c,"获取部门成功",menus)
	}
	return
}


// @Summary 添加部门
// @Description 添加部门
// @Tags 部门管理
// @accept json
// @Produce  json
// @Param Authorization header   string true "token"
// @Param body body  dto.SysDeptInput true "部门信息"
// @Success 200 {object}  middleware.Response{result=bool} "success"
// @Router /api/system/dept/add [post]
func CreateDept(c *gin.Context) {
	var dept dto.SysDept
	err := c.ShouldBindJSON(&dept)
	if err != nil {
		middleware.ResponseFail(c,201,err.Error())
		return
	}
		err = service.SaveDept(dept)

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
	err = service.SaveDept(dept)

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
