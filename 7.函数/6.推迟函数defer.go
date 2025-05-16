// @Author: MoYi
// @Date: 2025/5/16
// @Time: 19:39
package main

import (
	"fmt"
)

// 当函数执行到最后的时候  defer将会倒序执行
func main() {
	//f("1")
	//fmt.Println("2")
	//defer f("3")
	//fmt.Println("4")
	//defer f("5")
	//fmt.Println("6")
	//defer f("7")
	//fmt.Println("8")
	//defer f("9")
	a := 10
	fmt.Println("a=", a)
	defer f(a) // 参数已经传递进去了，在最后执行。
	a++
	fmt.Println("end a=", a)
}

func f(s int) {
	fmt.Println("函数里面的a", s)
}
