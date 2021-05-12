package service

import (
	//"bytes"
	"fmt"
	"gin-vben-admin/common"
	"gin-vben-admin/common/utils"
	"gin-vben-admin/dao"
	"gin-vben-admin/dto"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

func GetLogList(q dto.QuerySysOpLog) (err error, total int, logList []dto.SysOpLog) {
	var  OpLoglist []dao.SysOpLog
	table := common.DB.Model(&OpLoglist)
	if q.Type  == "1"{
		table = table.Where( "business_type in (?)", []int{10,11,12})
	}
	if q.Type  == "2"{
		table = table.Where( "business_type in (?)", []int{1,2,3,4})
	}

	if q.Method != ""{
		table = table.Where( "request_method = (?)",q.Method )
	}

	if q.OperName != ""{
		table = table.Where( "oper_name = (?)",q.OperName )
	}
	if q.IpAddr != ""{
		table = table.Where( "oper_ip = (?)",q.IpAddr )
	}

	if q.StartTime != ""{
		table = table.Where( "oper_time >= (?)",q.StartTime )
	}
	if q.EndTime != ""{
		table = table.Where( "oper_time <= (?)",q.EndTime )
	}

	table=table.Count(&total)
	if q.Page != "" && q.PageSize != "" {
	table = table.Limit(utils.StrToInt(q.PageSize)).Offset((utils.StrToInt(q.Page) - 1) * utils.StrToInt(q.PageSize))
	}
	if err = table.Order("id desc").Find(&logList).Error; err != nil {
			return err, 0, nil
	}


	return
}


func CreatOpLog(c *gin.Context, Username string, latencyTime time.Duration,Remark string ) {
	sysOperLog := dao.SysOpLog{}

	// 请求方式
	reqMethod := c.Request.Method
	// 请求路由
	reqUri := c.Request.RequestURI
	// 状态码
	statusCode := c.Writer.Status()

	ua := user_agent.New(c.Request.UserAgent())
	browserName, browserVersion := ua.Browser()
	sysOperLog.Browser = browserName + " " + browserVersion
	sysOperLog.Os = ua.OS()
	sysOperLog.Platform = ua.Platform()
	clientIP:=c.ClientIP()
	sysOperLog.Remark = Remark
	sysOperLog.OperIp = clientIP
	sysOperLog.OperLocation = utils.GetLocation(clientIP).(string)
	sysOperLog.Status = strconv.Itoa(statusCode)
	sysOperLog.OperName =Username
	sysOperLog.RequestMethod = reqMethod
	sysOperLog.OperUrl = reqUri
	if reqUri == "/api/login" {
		sysOperLog.BusinessType = "10"
		sysOperLog.Title = "用户登录"
		sysOperLog.OperName = Username
	} else if strings.Contains(reqUri, "/api/logout") {
		sysOperLog.Title = "用户登出"
		sysOperLog.BusinessType = "11"
	} else if strings.Contains(reqUri, "/api/system/base/captcha") {
		sysOperLog.BusinessType = "12"
		sysOperLog.Title = "验证码"
	} else {
		if reqMethod == "GET" {
			sysOperLog.BusinessType = "1"
		}else if reqMethod == "POST" {
			sysOperLog.BusinessType = "2"
		} else if reqMethod == "PUT" {
			sysOperLog.BusinessType = "3"
		} else if reqMethod == "DELETE" {
			sysOperLog.BusinessType = "4"
		}
	}
	payload, exist := c.Get("payload")
	if exist {
		sysOperLog.OperParam = payload.(string)
	}



	sysOperLog.CreateBy = Username
	sysOperLog.OperTime = time.Now()
	sysOperLog.LatencyTime = (latencyTime).String()
	if c.Err() == nil {
		sysOperLog.Status = "1"
	} else {
		sysOperLog.Status = "0"
	}

	loginInfo:=fmt.Sprintf("%3d %13v %15s %s %s ",
		statusCode,
		latencyTime,
		clientIP,
		reqMethod,
		reqUri)

	if err := common.DB.Save(&sysOperLog).Error; err != nil{
		common.LOG.Error("创建操作日志失败", zap.Any("err", err))
	}
	common.LOG.Info(loginInfo)
}

func BatchDeleteOperLog(ids []int) (err error) {
	err = common.DB.Unscoped().Model(&dao.SysOpLog{}).Where("id in (?)", ids).Delete(&dao.SysOpLog{}).Error
	return
}

func DeleteOperLog(id int) (err error) {
	err = common.DB.Unscoped().Model(&dao.SysOpLog{}).Where("id = (?)", id).Delete(&dao.SysOpLog{}).Error
	return
}




