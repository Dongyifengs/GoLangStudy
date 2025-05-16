// @Author: MoYi
// @Date: 2025/5/17
// @Time: 00:03
package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// Gin加载
func ginMain() {
	// 加载GinServer
	ginServer := gin.Default()

	// 登录API
	ginServer.POST("/login", func(context *gin.Context) {
		// 获取账号密码
		// 获取请求体（字节类型）
		data, _ := context.GetRawData()
		// 定义一个空的JSON对象用于接收请求
		var request UserLoginRequest
		// 将请求体数据存入JSON对象中
		_ = json.Unmarshal(data, &request)
		// 从其中取出需要查询的参数
		username := request.Username
		password := request.Password
		// 处理
		if username == "moyi" && password == "123" {
			context.JSON(http.StatusOK, gin.H{
				"username": username,
				"password": password,
				"msg":      "这是一个正确密码。",
			})
		} else if username == "" || password == "" {
			context.JSON(http.StatusOK, gin.H{
				"username": username,
				"password": password,
				"msg":      "账号密码请勿为空！",
			})
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"username": username,
				"password": password,
				"msg":      "账号密码不正确！",
			})
		}
	})

	ginServer.Run(":8000")

}

// 主入口
func main() {
	ginMain()
}
