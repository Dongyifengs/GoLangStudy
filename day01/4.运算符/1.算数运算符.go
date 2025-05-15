// @Author: MoYi
// @Date: 2025/5/15
// @Time: 16:09
package main

import "fmt"

// 基础的算数运算符
// 	+ 	相加
// 	- 	相减
// 	* 	相乘
// 	/ 	相除
// 	% 	求余
// 	++	自增
// 	--	自减

func main() {
	var a int = 10
	var b int = 3
	fmt.Printf("a = %d\nb = %d\n", a, b)
	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b)
	fmt.Println("a / b = ", a/b)
	fmt.Println("a % b = ", a%b)
	a++
	b--
	fmt.Println("a++ = ", a)
	fmt.Println("a-- = ", b)
}
