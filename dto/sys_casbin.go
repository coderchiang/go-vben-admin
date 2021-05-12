package dto

// Casbin info structure
type SysCasbin struct {
	RoleId string            `json:"role_id"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
