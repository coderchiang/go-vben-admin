package dto

type SysRole struct {
	ID string                  `json:"id" `
	RoleName string             ` json:"roleName" binding:"required"`
	RoleValue string             `json:"roleValue"`
	Remark string             `json:"remark"`
	Status string             `json:"status"`
	OrderNo string               `json:"orderNo"`
	CreatedAt       string   `json:"createTime"`
	Menu       []string   `json:"menu"`
}

type QuerySysRole struct {
	ID string                  `form:"id" json:"id" `
	RoleName string             `form:"roleName" json:"roleName"`
	Status string             `form:"status" json:"status"`
	PageSize     string    `form:"pageSize" json:"pageSize"`
	Page      string    `form:"page" json:"page"`
}

