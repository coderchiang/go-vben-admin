package dto

type SysMenu struct {
	ID     string             `json:"id" `
	Pid     string             `json:"pid" `
	ParentName string        `json:"pMenuName"`
	Name string             `json:"name"`
	MenuName string             `json:"menuName"   binding:"required"`
	Path string             `json:"path" `
	Sort int                `json:"orderNo" `
	CreateTime  string     `json:"createTime" `
	MenuType  string                `json:"type"`
	Status  string              `json:"status"`
	Component string        `json:"component" `
	Icon string             `json:"icon" `
	ApiPath string         `json:"apiPath"`
	ApiMethod string       `json:"apiMethod"`
	Keepalive string            `json:"keepalive"`
	IsExt    string              `json:"is_ext"`
	Meta Meta               `json:"meta" `
	Children []*SysMenu     `json:"children"`

}
type  Meta struct {
	Affix bool             `json:"affix" `
	Title      string      `json:"title"`
	Icon string             `json:"icon" `
	HideTab bool             `json:"hideTab" `
	HideChildrenInMenu bool             `json:"hideChildrenInMenu" `
	IgnoreKeepAlive bool      `json:"ignoreKeepAlive" `
}

type QuerySysMenu struct {
	ID string                  `form:"id" json:"id" `
	Name string             `form:"menuName" json:"menuName"`
	Status string             `form:"status" json:"status"`
	PageSize     string    `form:"pageSize" json:"pageSize"`
	Page      string    `form:"page" json:"page"`
}