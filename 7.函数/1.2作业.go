// @Author: MoYi
// @Date: 2025/5/16
// @Time: 18:51
package main

import "fmt"

func main() {
	fmt.Println(maxTowNum(10, 20))
}

// Max 两个数字比大小
func maxTowNum(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
