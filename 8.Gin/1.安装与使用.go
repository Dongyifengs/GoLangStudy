// @Author: MoYi
// @Date: 2025/5/16
// @Time: 22:08
package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"log"
	"net/http"
)

// 自定义拦截器 -> 中间件
func myHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 通过定义中间件，设置的值，在后续处理只要调用了这个中间件都可以拿到这里的参数
		context.Set("userSession", "userid-1")
		//if xx.xx {
		//	context.Next() // 放行
		//} else {
		//	context.Abort() // 阻止
		//}
		context.Next() // 放行
	}
}

func main() {

	// 创建一个服务
	ginServer := gin.Default()
	ginServer.Use(favicon.New("./img/favicon2.ico"))

	// 加载静态页面
	ginServer.LoadHTMLGlob("templates/*")

	// 加载资源文件
	ginServer.Static("/static", "./static")

	// 相应一个页面给前端
	ginServer.GET("/index", func(context *gin.Context) {
		// context.JSON() son数据
		context.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "这是Go后台转递的数据",
		})
	})

	// url?userid=xxx&username=moyi
	ginServer.GET("/user/info", myHandler(), func(context *gin.Context) {

		// 取出中间件的值
		usersession := context.MustGet("userSession").(string)
		log.Println("------->", usersession)

		userid := context.Query("userid")
		username := context.Query("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	// /user/info/1/moyi
	ginServer.GET("/user/info/:userid/:username", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	// 后面的应用 - 基础知识 + 了解web开发
	// 前端给后端传递JSON
	ginServer.POST("/json", func(context *gin.Context) {
		// request.body
		// []byte
		b, _ := context.GetRawData()

		var m map[string]interface{}
		// 包装为json对象
		_ = json.Unmarshal(b, &m)
		context.JSON(http.StatusOK, m)

	})

	// 表单处理的情况
	ginServer.POST("/user/add", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")

		// 数据判断可以在这里做

		context.JSON(http.StatusOK, gin.H{
			"msg":      "OK",
			"username": username,
			"password": password,
		})
	})

	// 路由
	ginServer.GET("/test", func(context *gin.Context) {
		// 重定向
		context.Redirect(http.StatusMovedPermanently, "https://wzpmc.cn:83")
	})

	// 404
	ginServer.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", nil)
	})

	// 路由组
	userGroup := ginServer.Group("/user")
	{
		userGroup.GET("/add")
		userGroup.GET("/login")
		userGroup.GET("/logout")
	}
	orderGroup := ginServer.Group("/order")
	{
		orderGroup.GET("/add")
		orderGroup.GET("/delte")
	}

	// 服务器端口
	ginServer.Run(":8000")
}
