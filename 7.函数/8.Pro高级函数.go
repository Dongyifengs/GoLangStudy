// @Author: MoYi
// @Date: 2025/5/16
// @Time: 19:54
package main

import "fmt"

func main() {
	f111()
	f2 := f111
	f2()

	// miming函数
	f3 := func() {
		fmt.Println("我是f3函数")
	}
	f3()

	// miming函数自调用
	func(a, b int) {
		fmt.Println(a, b)
		fmt.Println("我是f4函数")
	}(1, 2)

	r1 := func(a, b int) int {
		return a + b
		// fmt.Println("我是f5函数")
	}(1, 2)
	fmt.Println(r1)
}

func f111() {
	fmt.Println("我是f111函数")
}
