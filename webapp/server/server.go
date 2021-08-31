// @author: lls
// @date: 2021/7/15
// @desc:

package main

import (
	"github.com/gin-gonic/gin"
	"golearn/util"
	"net/http"
	"sync/atomic"
)

type Request map[string]interface{}

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

// 跨域访问：cross  origin resource share
func CrossHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Access-Control-Allow-Credentials", "true")
		context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,token,openid,opentoken")
		context.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE,UPDATE")
		context.Header("Access-Control-Allow-Origin", context.GetHeader("origin")) // 设置允许访问域
		if context.Request.Method == "OPTIONS" {
			context.JSON(http.StatusOK, gin.H{})
		}
		// 处理请求
		context.Next()
	}
}

var id int32

func main() {
	r := gin.Default()
	r.Use(CrossHandler())
	r.POST("/save_form", func(context *gin.Context) {
		atomic.AddInt32(&id, 1)
		context.JSON(http.StatusOK, Response{
			Status: 0,
			Msg:    "",
			Data: map[string]interface{}{
				"id": id,
			},
		})
	})
	util.MustNil(r.Run("0.0.0.0:2021")) // 监听并在 0.0.0.0:8080 上启动服务
}
