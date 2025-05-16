// @Author: MoYi
// @Date: 2025/5/16
// @Time: 18:41
package main

import "fmt"

func main() {
	// 函数的调用
	printInfo()
	myPrint("哈哈")
	myPrintNum(10086)
	fmt.Println(add2(2, 4))
	x, y := swap("你是", "Wzp")
	fmt.Println(x, y)
	fmt.Println(swap("我是", "MoYi"))
}

// 无参无返回值函数
func printInfo() {
	fmt.Println("printInfo")
}

// 有一个参数的函数
func myPrint(msg string) {
	fmt.Println(msg)
}

func myPrintNum(num int) {
	fmt.Println(num)
}

// 有两个参数的函数
// 有一个返回值的函数
func add2(a, b int) int {
	c := a + b
	return c
}

// 有多个返回值的函数
func swap(x, y string) (string, string) {
	return x, y
}
