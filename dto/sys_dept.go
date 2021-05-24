package dto

type SysDeptOutput struct {
	Items []*SysDept  `json:"items" `
	Total  int       `json:"total" `

}

type SysDept  struct {
	ID     string             `json:"id" `
	Pid     string             `json:"pid" `
	DeptName string        `json:"deptName" binding:"required"`
	OrderNo int                `json:"orderNo" `
	CreateTime  string     `json:"createTime" `
	Status  string              `json:"status"`
	Remark  string              `json:"remark"`
	ParentDept string            `json:"parentDept"`
	Children []*SysDept     `json:"children"`

}

type SysDeptInput struct {
	ID     string             `json:"id" `
	Pid     string             `json:"pid" `
	DeptName string        `json:"deptName" binding:"required"`
	OrderNo int                `json:"orderNo" `
	Status  string              `json:"status"`
	Remark  string              `json:"remark"`
	ParentDept string            `json:"parentDept"`
}


type QuerySysDept struct {
	ID string                  `form:"id" json:"id" `
	DeptName string             `form:"deptName" json:"deptName"`
	Status string             `form:"status" json:"status"`
	PageSize     string    `form:"pageSize" json:"pageSize"`
	Page      string    `form:"page" json:"page"`
}