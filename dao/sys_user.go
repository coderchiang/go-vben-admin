package dao

import "time"

type SysUser struct {
	ID        int      `gorm:"primary_key" json:"id"`
	Username  string    `json:"username" gorm:"unique_index;not null"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"password"`
	AvatarUrl string    `json:"avatar_url" gorm:"default:'static/upload/avatar/default.png'"`
	RoleId    int       `json:"role_id" `
	Status      int     `json:"status" gorm:"default:1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Dept      int       `json:"dept_id" `
	Phone      string       `json:"phone" `
	Email      string       `json:"email" `
	Remark      string       `json:"remark" `
	SysRole      SysRole   `gorm:"ForeignKey:RoleId" json:"role"`
}

