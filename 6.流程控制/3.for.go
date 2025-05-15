// @Author: MoYi
// @Date: 2025/5/15
// @Time: 18:52
package main

import "fmt"

func main() {
	// Study()
	// Test1()
	// Test2()
	BC()
}

func Study() {
	// for 条件的起始值，循环条件，控制变量自增或者自减
	for i := 1; i <= 10; i++ {
		fmt.Println("循环次数为", i)
	}

	//i := 1
	//for i <= 10 {
	//	fmt.Println("循环次数为", i)
	//	i++
	//}

	//i := 1
	//for {
	//	fmt.Println(i)
	//	i++
	//}

	// 计算1-10的和
	//var z int = 0
	//for i := 1; i <= 10; i++ {
	//	z = z + i
	//}
	//fmt.Println("和为", z)
}
func Test1() {
	/* 打印一个方阵
	* * * *
	* * * *
	* * * *
	* * * *
	 */

	var x int = 5
	var y int = 5

	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
func Test2() {

	for j := 1; j <= 9; j++ {
		for i := 1; i <= j; i++ {
			fmt.Printf("%d * %d = %d \t", i, j, i*j)
		}
		fmt.Println()
	}

}
func BC() {
	// break 直接中止循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("============================")

	// continue 结束单次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
