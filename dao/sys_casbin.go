package dao

import (
	"sync"
)

type CasbinRule struct {
	Ptype       string `json:"ptype" gorm:"column:p_type"`
	RoleId      string `json:"rolen_id" gorm:"column:v0"`
	Path        string `json:"path" gorm:"column:v1"`
	Method      string `json:"method" gorm:"column:v2"`
}


type Auth struct {
	Ch chan bool
	Wg *sync.WaitGroup
}
func (casbin CasbinRule)TableName() string  {
	return "casbin_rule"
}

