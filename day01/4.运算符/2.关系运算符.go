// @Author: MoYi
// @Date: 2025/5/15
// @Time: 16:18
package main

import "fmt"

func main() {

	var a int = 11
	var b int = 10

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

	// if -> 如果。。。结果
	if a > b {
		fmt.Println("a>b")
	} else {
		fmt.Println("a<b")
	}

}
