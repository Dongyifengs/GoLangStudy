// @Author: MoYi
// @Date: 2025/5/15
// @Time: 18:10
package main

import "fmt"

func main() {

	var x int
	var y float64

	// 定义了两个变量，想用键盘来录入这两个变量
	// fmt.Println()	// 打印并换行
	// fmt.Printf()		// 格式化输出
	// fmt.Print()		// 打印输出

	fmt.Println("请输入两个数：1、整数，2、浮点数")
	// 变量取 地址 & 变量 | 指针、地址来修改和操作变量
	// Scanln 阻塞等待你的键盘输入
	fmt.Scanln(&x, &y)
	fmt.Println("X:", x)
	fmt.Println("Y:", y)

	// fmt.Scan()		// 接受输入 Scan |
	// fmt.Scanf()		// 接收输入 格式化输入 作业
	// fmt.Scanln()		// 接收输入 作业
}
