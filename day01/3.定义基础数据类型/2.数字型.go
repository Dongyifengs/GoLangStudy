package main

import "fmt"

func main() {
	// 定义一个整形
	var age int = 18
	fmt.Printf("%T,%d", age, age)

	// 定义一个浮点型
	// 默认是六位小数打印 3.140000
	var money float64 = 3.14
	fmt.Printf("%T,%f.1f", money, money)
	fmt.Printf("%T,%.1f", money, money) // 会四舍五入,保留一位小数

	// World
	var num1 float32 = -123.0000901
	var num2 float64 = -123.0000901
	fmt.Println("num1 = ", num1, "num2 = ", num2)
}
