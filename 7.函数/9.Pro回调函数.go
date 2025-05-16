// @Author: MoYi
// @Date: 2025/5/16
// @Time: 20:02
package main

import "fmt"

func main() {
	r1 := test1(1, 2)
	fmt.Println(r1)

	r2 := oper(3, 4, test1)
	fmt.Println(r2)

	r3 := oper(3, 4, sub)
	fmt.Println(r3)

	r4 := oper(8, 2, func(i int, i2 int) int {
		if i2 == 0 {
			fmt.Println("除数不能为0")
			return 0
		}
		return i / i2
	})
	fmt.Println(r4)
}

// 高阶函数,可以接收一个函数作为参数
func oper(a, b int, fun func(int, int) int) int {
	r := fun(a, b)
	return r
}

func test1(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
