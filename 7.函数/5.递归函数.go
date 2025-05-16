// @Author: MoYi
// @Date: 2025/5/16
// @Time: 19:20
package main

import "fmt"

func main() {
	fmt.Println(getSum(1))
}

func getSum(nums int) int {
	Sums := 0
	if nums == 1 {
		Sums = nums
		return Sums
	} else {
		for i := 0; i <= nums; i++ {
			Sums = Sums + i
		}
		return Sums
	}

}
