package main

import "fmt"

func main() {

	var str string
	str = "Hello,Go"
	fmt.Printf("%T,%s\t", str, str)

	// 单引号
	v1 := 'A'
	v2 := "A"
	fmt.Printf("%T,%d\t", v1, v1) // int32,65
	fmt.Printf("%T,%s\t", v2, v2) // str,A

	// 字符串连接
	fmt.Printf("Hello" + "World")

	// 转义 \" 可以直接转义出来
	fmt.Printf("Hello\"World") // 打印”
	fmt.Printf("Hello\nWorld") // n -> 换行
	fmt.Printf("Hello\tWorld") // t -> 制表符

}
