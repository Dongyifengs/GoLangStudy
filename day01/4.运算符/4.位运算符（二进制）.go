// @Author: MoYi
// @Date: 2025/5/15
// @Time: 16:18
package main

import "fmt"

// 针对于二进制
func main() {
	var a uint = 60
	var b uint = 13
	var c uint = 0

	c = a & b // 位运算符
	fmt.Printf("当C = A & B时，结果为:%d,二进制:%b\n", c, c)

	c = a | b // 位运算符
	fmt.Printf("当C = A | B时，结果为:%d,二进制:%b\n", c, c)

	c = a ^ b // 异或运算符
	fmt.Printf("当C = A ^ B时，结果为:%d,二进制:%b\n", c, c)

	c = a &^ b // 位清空
	fmt.Printf("当C = A &^ B时，结果为:%d,二进制:%b\n", c, c)

	c = a << b // 左移运算符
	fmt.Printf("当C = A << B时，结果为:%d,二进制:%b\n", c, c)

	c = a >> b // 右移运算符
	fmt.Printf("当C = A >> B时，结果为:%d,二进制:%b\n", c, c)
}
