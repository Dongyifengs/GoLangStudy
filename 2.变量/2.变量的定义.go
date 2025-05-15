package main

import "fmt"

func main() {
	//TIP 定义变量使用var关键字 如果没有赋值，则是这个类型的默认值
	// string 默认值为 空
	// int 默认值为 0
	var id int = 1
	var a, b, c *int
	var d, e, f int = 1, 2, 3
	var (
		name    string
		age     int
		address string
	)
	fmt.Println(name, age, address, id)
	fmt.Println(a, b, c, d, e, f)
}
