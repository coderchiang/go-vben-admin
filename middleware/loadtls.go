package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

//HTTPS配置步骤:
//
//首先在阿里云搞定ICP域名备案
//添加一个子域名
//给子域名申请免费 SSL 证书, 然后下载证书对应的 pem 和 key 文件.
//用 GIN 框架添加一个 github.com/unrolled/secure 中间件就可以了.
// 用https把这个中间件在router里面use一下就好


func LoadTls() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:443",
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			//如果出现错误，请不要继续。
			fmt.Println(err)
			return
		}
		// 继续往下处理
		c.Next()
	}
}
