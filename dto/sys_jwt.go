package dto

import "github.com/dgrijalva/jwt-go"

// 载荷，可以加一些自己需要的信息
type Claims struct {
	UserID        string    `json:"user_id"`
	Username  string `json:"username"`
	RoleId string `json:"role_id"`
	jwt.StandardClaims
}
