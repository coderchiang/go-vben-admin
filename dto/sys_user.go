package dto

type PassWord struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"pwd" json:"pwd" binding:"required"`
	NewPassword string `form:"new_pwd" json:"new_pwd" binding:"required"`
}

type UserLogin struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	//Captcha   string `json:"captcha" binding:"required"`
	//CaptchaId string `json:"captchaId" binding:"required"`
}

type LoginSucc struct {
	Token     string    `json:"token"`
	UserInfo  UserInfoOut   `json:"userinfo"`
}

type UserInfoOut struct{
	ID        string    `json:"id"`
	Username  string    `json:"username"  binding:"required"`
	Nickname  string    `json:"nickname" `
	RoleId    string       `json:"role_id" `
	RoleName    string       `json:"roleName" `
	Phone      string       `json:"phone" `
	Email      string       `json:"email" `
	Dept      string       `json:"deptName" `
	Status     int      `json:"status" `
	Remark      string       `json:"remark" `
	CreateTime      string       `json:"createTime" `
}

type UserInfoIn struct{
	ID        string    `json:"id"`
	Username  string    `json:"username"  binding:"required"`
	Nickname  string    `json:"nickname" `
	Password string     `form:"pwd" json:"pwd"`
	AvatarUrl string    `json:"avatar_url"`
	RoleId    string       `json:"role_id" `
	RoleName    string       `json:"roleName" `
	Dept      string       `json:"deptName" `
	Phone      string       `json:"phone" `
	Email      string       `json:"email" `
	Remark      string       `json:"remark" `
	Status     string      `json:"status" `
}

type QueryUser struct {
	ID       string `form:"id" json:"id" `
	Username string `form:"username" json:"username"`
	Phone    string `form:"phone" json:"phone" `
	DeptId   string `form:"deptId" json:"deptId" `
	PageSize string `form:"pageSize" json:"pageSize"`
	Page     string `form:"page" json:"page"`
}

