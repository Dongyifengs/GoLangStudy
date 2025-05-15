package main

import "fmt"

func main() {

	//TIP 个人感觉类似Python，直接使用 目标类型() 转换即可

	a := 5.0
	b := int(a)

	fmt.Printf("%T,%f", a, a)
	fmt.Printf("%T,%d", b, b)
}
