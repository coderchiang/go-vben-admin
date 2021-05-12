package dao

import "time"

type SysDept struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Pid     int             `json:"pid" `
	DeptName string        `json:"deptName" binding:"required"`
	OrderNo int                `json:"orderNo" `
	Level int                `json:"level" `
	Status  int              `json:"status"`
	Remark  string              `json:"remark"`
}
