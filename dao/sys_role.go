package dao

import "time"

type SysRole struct {
	ID int                  `json:"role_id" gorm:"not null;primary_key"`
	RoleName string             `gorm:"unique_index;not null" json:"role_name"`
	RoleValue string             `json:"role_value"`
	Remark string             `json:"remark"`
	Status string             `json:"status"`
	OrderNo int                `json:"order_no"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"`
	SysMenu  []*SysMenu      `gorm:"many2many:sys_role_menu"`
}

func (SysRole) TableName() string {
	return "sys_role"
}

