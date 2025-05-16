// @Author: MoYi
// @Date: 2025/5/16
// @Time: 19:15
package main

import "fmt"

// 全局变量
var num int = 100

func main() {
	temp := 100

	// if 语句 for语句定义的一次性变量是局部变量
	if b := 1; b <= 10 {
		// 语句内的局部变量
		temp := 50
		fmt.Println(temp) // 局部变量，就近原则
		fmt.Println(b)
	}
	num := 20
	fmt.Println(temp)
	fmt.Println(num)
	f1()
	f2()
}

func f1() {
	a := 1
	num := 30
	fmt.Println(a)
	fmt.Println(num)
}

func f2() {
	num := 40
	// fmt.Println(a) // 不能使用其他函数使用的变量
	fmt.Println(num)
}
