package system

import (
	"gin-vben-admin/common"
	"gin-vben-admin/common/utils"
	"gin-vben-admin/dao"
	"gin-vben-admin/dto"
	"gin-vben-admin/middleware"
	"gin-vben-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetOperLogList(c *gin.Context)  {
	var q dto.QuerySysOpLog
	if err := c.ShouldBind(&q); err != nil {
		middleware.ResponseFail(c, 201, err.Error())
		return
	}
	err, total, logList := service.GetLogList(q)
	if err != nil {
		middleware.ResponseFail(c, 202, err.Error())
		return
	}
	middleware.ResponseSucc(c, "获取用户列表成功", map[string]interface{}{
		"items":logList,
		"total": total,
	})
	return
}


func BatchDeleteLog(c *gin.Context) {
	var idsStr dto.QuerySysOpLog
	err := c.ShouldBind(&idsStr)
	if err!=nil {
		middleware.ResponseFail(c,  201,err.Error())
		return
	}
	var data dao.SysOpLog
	data.UpdateBy = string(middleware.GetClaims(c).UserID)
	var ids []int
	for _, id := range idsStr.Ids {
		x:= utils.StrToInt(id)
		ids = append(ids, x)
	}
	err = service.BatchDeleteOperLog(ids)
	if err != nil {
		middleware.ResponseFail(c, 204, "批量删除失败")
		common.LOG.Error("批量删除失败",zap.Any("err", err))
		return
	}
	middleware.ResponseSucc(c, "删除成功",true)
	return
}


func DeleteLog(c *gin.Context)  {

	var idsStr dto.QuerySysOpLog
	err := c.ShouldBind(&idsStr)
	if err!=nil {
		middleware.ResponseFail(c,  201,err.Error())
		return
	}

	if err = service.DeleteOperLog(utils.StrToInt(idsStr.ID)); err != nil{
		middleware.ResponseFail(c, 202, "删除失败")
		common.LOG.Error("删除失败",zap.Any("err", err))
		return
	}
	middleware.ResponseSucc(c, "删除成功",true)
	return
}