package dao

import "time"

type SysOpLog struct {
	ID        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"` //日志编码
	Title         string    `json:"title" gorm:"size(255);"`                  //操作模块
	BusinessType  string    `json:"businessType" gorm:"type:varchar(128);"`   //操作类型
	RequestMethod string    `json:"requestMethod" gorm:"type:varchar(128);"` //请求方式
	OperatorType  string    `json:"operatorType" gorm:"type:varchar(128);"`  //操作类型
	OperName      string    `json:"operName" gorm:"type:varchar(128);"`      //操作者
	DeptName      string    `json:"deptName" gorm:"type:varchar(128);"`      //部门名称
	OperUrl       string    `json:"operUrl" gorm:"type:varchar(255);"`       //访问地址
	OperIp        string    `json:"operIp" gorm:"type:varchar(128);"`        //客户端ip
	OperLocation  string    `json:"operLocation" gorm:"type:varchar(128);"`  //访问位置
	OperParam     string    `json:"operParam" gorm:"type:varchar(255);"`     //请求参数
	Status        string    `json:"status" gorm:"type:int(1);"`              //操作状态
	OperTime      time.Time `json:"operTime" gorm:"type:datetime;"`         //操作时间
	JsonResult    string    `json:"jsonResult" gorm:"type:varchar(255);"`    //返回数据
	CreateBy      string    `json:"createBy" gorm:"type:varchar(128);"`      //创建人
	UpdateBy      string    `json:"updateBy" gorm:"type:varchar(128);"`      //更新者
	Browser       string     `json:"browser" gorm:"type:varchar(128);"`
	Os            string     `json:"os" gorm:"type:varchar(128);"`
	Platform      string     `json:"platform" gorm:"type:varchar(128);"`
	Remark        string    `json:"remark" gorm:"type:varchar(255);"`        //备注
	LatencyTime   string    `json:"latencyime" gorm:"type:varchar(128);"`    //耗时
	CreatedAt time.Time `json:"createdAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

func (SysOpLog) TableName() string {
	return "sys_op_log"
}

