// @Author: MoYi
// @Date: 2025/5/15
// @Time: 16:19
package main

import "fmt"

func main() {
	var a int = 21
	var b int

	// =	赋值
	b = a
	fmt.Println(b) // 21 = 0 + 21

	// +=	b = a + b
	b += a
	fmt.Println(b) // 42 = 21 + 21

	// -= *= /= %= <<= >>= &= ^= |=
}
