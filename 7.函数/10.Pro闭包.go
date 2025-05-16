// @Author: MoYi
// @Date: 2025/5/16
// @Time: 20:12
package main

import "fmt"

func store() func() int {
	i := 1
	add := func() int {
		i++
		return i
	}
	return add
}

func main() {
	myStore := store()
	fmt.Println(myStore())
	fmt.Println(myStore())
}
