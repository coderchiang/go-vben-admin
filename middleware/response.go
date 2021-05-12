package middleware

import (
	"github.com/gin-gonic/gin"
)
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Result interface{} `json:"result"`
}
func ResponseSucc(c *gin.Context, msg string,data interface{})  {
	c.JSON(200,Response{Code:200,Msg:msg,Result: data})
	return
}

func ResponseFail (c *gin.Context,code int, msg string)  {
	c.JSON(code, Response{Code:code,Msg:msg,Result: nil})
	return
}
