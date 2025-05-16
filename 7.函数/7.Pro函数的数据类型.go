// @Author: MoYi
// @Date: 2025/5/16
// @Time: 19:45
package main

import "fmt"

// func() 本事呢就是一个数据类型
func main() {
	// f11 如果不加括号 函数就是一个变量
	// f11() 如果加括号里那就成了函数的调用
	fmt.Printf("%T\n", f11)     // func()	| func(int, int) | func(iny, int) int
	fmt.Printf("%T\n", 10)      // int
	fmt.Printf("%T\n", "hello") // string

	// 定义函数类型的变量
	var f5 func(int, int)
	f5 = f11 // 引用类型
	fmt.Println(f5)
	fmt.Println(f11)
	f5(1, 2)

}

func f11(a, b int) {
	fmt.Println(a, b)
}
