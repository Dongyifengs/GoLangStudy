// @Author: MoYi
// @Date: 2025/5/15
// @Time: 16:18
package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false

	// 都为true结果才为true
	if a == true && b == false {
		fmt.Println(a && b)
	}

	// 其中一个满足true即可
	if a || b {
		fmt.Println(a || b)
	}

	// 取反
	if !(a || b) {
		fmt.Println(!(a || b)) // 不会输出
	}

}
