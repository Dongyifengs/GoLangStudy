// @Author: MoYi
// @Date: 2025/5/16
// @Time: 18:57
package main

import "fmt"

func main() {
	fmt.Println("Test")
	getSun(1231, 312, 31, 31, 231, 23, 12, 31, 231, 32, 123, 12, 31, 3, 123, 1)
}

// ,,, 可变参数
func getSun(nums ...int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		fmt.Println(i)
		sum = sum + nums[i]
	}
	fmt.Println("sum:", sum)
	return sum
}
