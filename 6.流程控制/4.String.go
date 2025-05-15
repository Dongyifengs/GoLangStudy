// @Author: MoYi
// @Date: 2025/5/15
// @Time: 19:18
package main

import "fmt"

func main() {
	str := "helloworld"
	fmt.Println(str)

	// 获取字符串长度
	fmt.Println("字符串的长度为：", len(str))

	// 获取指定字节
	fmt.Println("字节打印", str[0]) // ACSII
	fmt.Printf("%c\n", str[2])

	// for
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}

	// for range循环，遍历数组，切片
	// 返回下标和对应的值，使用这个值就可以了
	for i, v := range str {
		fmt.Print(i)
		fmt.Printf("%c", v)
	}

	// String字符串不可以修改
	// str[2] = 'A'
	fmt.Println(str[2])
}
