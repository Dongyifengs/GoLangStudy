// @Author: MoYi
// @Date: 2025/5/17
// @Time: 21:25
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// GinMain 作为入口
func GinMain() {

	// 引用Gin
	ginServer := gin.Default()

	// 渲染网页
	ginServer.LoadHTMLGlob("template/*")
	ginServer.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "add.html", gin.H{
			"Title": "Main",
		})
	})

	// 制作增加接口
	ginServer.POST("/add", func(context *gin.Context) {
		// 获取页面提交的url
		urlText := context.PostForm("inputUrl")

		// 数据返回通知
		context.JSON(http.StatusOK, gin.H{
			"msg": "成功传入服务器后端",
			"url": urlText,
		})

	})

	// 开启端口运行框架
	err := ginServer.Run(":8000")
	if err != nil {
		return
	}

}

// UrlData 自定义类
type UrlData struct {
	OldUrl string `json:"old_url"`
	NewUrl string `json:"new_url"`
}

// ReadJsonFile 读取Json的文件内容
// filePath		:	文件路径
// []UrlData	:	返回UrlData
func ReadJsonFile(filePath string) []UrlData {
	// 定义data为UrlData的接收内容
	var data []UrlData
	// 读取文件
	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取Json文件失败！", err)
		return data
	}

	// 解析数据
	err1 := json.Unmarshal(jsonFile, &data)
	if err1 != nil {
		fmt.Println("解析Json文件失败！")
		return data
	}

	return data
}

// WriteJsonFile 写入Json文件
// filePath		:	文件路径
// data			:	写入数据
func WriteJsonFile(filePath string, data []UrlData) error {
	// 打开Json文件
	jsonFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("打开Json文件失败了！", err)
		return err
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println("未能正常关闭Json文件！", err)
		}
	}(jsonFile)

	// 格式化Json输出
	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return err

}

func main() {
	// GinMain()
	// 写入的数据
	arr := []UrlData{
		{
			"111",
			"2223",
		},
	}
	fmt.Println(WriteJsonFile("../json/json.json", arr))
	fmt.Println(ReadJsonFile("../json/json.json"))
}
