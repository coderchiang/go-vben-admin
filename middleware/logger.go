package middleware

import (
	"bytes"
	"gin-vben-admin/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		//

		// 把request的内容读取出来
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
			c.Set("payload",string(bodyBytes))
		}
		// 把刚刚读出来的再写进去
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))


		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		//如果username为空，就是登录逻辑
		userName :=GetClaims(c).Username


		if  userName!=""&&c.Request.Method != "OPTIONS" &&c.Request.Method!="GET"{
			//fmt.Println(c.Request.Method)
			go service.CreatOpLog(c,userName, latencyTime,"Operation log")
		}

	}
}
