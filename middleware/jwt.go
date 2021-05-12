package middleware

import (
	"errors"
	"gin-vben-admin/common"
	"gin-vben-admin/dto"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)



// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			ResponseFail(c,250,"请求未携带token，无权限访问")
			c.Abort()
			return
		}

		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {

			if err == TokenExpired {
				ResponseFail(c,251,"token 已经过期，请重新登录！")
				c.Abort()
				return
			}
			ResponseFail(c,252, err.Error())
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
		c.Next()
	}
}
// JWT 签名结构
type JWT struct {
	SigningKey []byte
}
// 一些常量
var (
	TokenExpired           = errors.New("Token is expired")
	TokenNotValidYet       = errors.New("Token not active yet")
	TokenMalformed         = errors.New("That's not even a token")
	TokenInvalid           = errors.New("Couldn't handle this token:")
	TokenExpireAt        int64  = 60*60*2		// 默认过期时间2小时过期
	Issuer                 = "chrischiang"
)



// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{[]byte(common.CONFIG.Jwt.SigningKey)}
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims dto.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*dto.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dto.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*dto.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &dto.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*dto.Claims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix() // 默认token有效期为1个小时
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

// 生成令牌
func GenerateToken(user  *dto.UserInfoOut) (token string, msg string, ok bool) {
	j := &JWT{[]byte(common.CONFIG.Jwt.SigningKey)}
	claims := dto.Claims{
		user.ID,
		user.Username,
		user.RoleId,

		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),          // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + TokenExpireAt), // 过期时间 一小时
			Issuer:    Issuer,                                   //签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		common.LOG.Error("创建Token失败", zap.Any("err", err))
		return token, "创建token失败", false
	} else {
		return token, "登录成功", true
	}
}

func GetClaims(c *gin.Context) *dto.Claims {
	claims, exist := c.Get("claims")
	if !exist {
		return &dto.Claims{}
	}
	return claims.(*dto.Claims)
}
