// @Author: MoYi
// @Date: 2025/5/15
// @Time: 18:41
package main

import "fmt"

// switch语句
func main() {
	switchS()
	// fallthroughS()
}

// 默认switch语句
func switchS() {
	var score int = 90

	// case 匹配
	switch score {
	case 90:
		fmt.Println("A")
		fallthrough // 很少用
	case 80:
		if score == 90 {
			break
		}
		fmt.Println("B")
		fallthrough
	case 50, 60, 70:
		fmt.Println("C")
		fallthrough
	default:
		fmt.Println("D")
	}

	// 啥都没有默认是布尔值，为True
	switch {
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	default:
		fmt.Println("其他")
	}
}
func fallthroughS() {
	a := false
	switch a {
	case false:
		fmt.Println("1、case条件为false")
		fallthrough // case穿透，不管下一个条件满不满足，都会执行
	case true:
		if a == false {
			break // 中止case穿透
		}
		fmt.Println("2、case条件为true")
	}

}
