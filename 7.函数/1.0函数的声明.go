// @Author: MoYi
// @Date: 2025/5/16
// @Time: 18:36
package main

import "fmt"

func main() {
	fmt.Println("HelloWorld")
	fmt.Println(add(1, 2))
}

/*
	func 函数名 （参数，参数 。。。） 函数调用后的返回值 {
		函数体 : 执行一段代码
		return 返回结果
*/
func add(a, b int) int {
	c := a + b
	return c
}
