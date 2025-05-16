// @Author: MoYi
// @Date: 2025/5/16
// @Time: 19:11
package main

import "fmt"

// 引用传递
func main() {
	// 切片，可以扩容的数组
	s1 := []int{1, 2, 3, 4}
	fmt.Println("默认的数据S1", s1)
	update2(s1)
	fmt.Println("修改后的数据S1", s1)
}

func update2(s2 []int) {
	fmt.Println("传递的数据s2", s2)
	s2[0] = 100
	fmt.Println("修改后的数据s2", s2)
}
