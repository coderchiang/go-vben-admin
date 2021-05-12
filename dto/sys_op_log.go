package dto

type QuerySysOpLog struct {
	ID     string            `form:"id" json:"id"`
	Ids []string `form:"ids" json:"ids"`
	PageSize     string    `form:"pageSize" json:"pageSize"`
	Page      string    `form:"page" json:"page"`
	Type         string `form:"type" json:"type"`
	Method       string `form:"method" json:"method"`
	IpAddr       string `form:"ipaddr" json:"ipaddr"`
	OperName     string `form:"operName" json:"operName"`
	StartTime       string `form:"startTime" json:"startTime"`
	EndTime       string `form:"endTime" json:"endTime"`
}


type SysOpLog struct {
	ID     string            `form:"id" json:"id"`
	RequestMethod string    `json:"method"` //请求方式
	OperatorType  string    `json:"type""`  //操作类型
	OperName      string    `json:"operName" `      //操作者
	OperUrl       string    `json:"operUrl"`       //访问地址
	OperIp        string    `json:"operIp" `        //客户端ip
	OperLocation  string    `json:"operLocation" `  //访问位置
	Status        string    `json:"status" `              //操作状态
	OperTime      string     `json:"operTime" `         //操作时间
	Browser       string     `json:"browser" `
	Os            string     `json:"os" `
	Platform      string     `json:"platform" `
	Remark        string    `json:"remark" `        //备注
	LatencyTime   string    `json:"latencyTime" `    //耗时
}

