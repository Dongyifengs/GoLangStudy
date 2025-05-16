// @Author: MoYi
// @Date: 2025/5/16
// @Time: 19:06
package main

import "fmt"

func main() {

	// 值传递
	// 定义一个数组 := [个数]类型
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	// 传递：拷贝arr
	update(arr)
	fmt.Println("调用后的数据", arr)

}

func update(arr2 [4]int) {
	fmt.Println("arr2接收的数据", arr2)
	arr2[0] = 100
	fmt.Println("arr2修改后的数据", arr2)
}
