// @Author: MoYi
// @Date: 2025/5/15
// @Time: 18:19
package main

import "fmt"

func main() {
	// ifStatement()
	ifNestedStatements()
}

// if...else... -> 如果...否则...
// if语句
func ifStatement() {
	var score int = 50
	if score >= 90 && score <= 100 {
		fmt.Println("A")
	} else if score >= 80 && score < 90 {
		fmt.Println("B")
	} else if score >= 70 && score < 80 {
		fmt.Println("C")
	} else if score >= 60 && score < 70 {
		fmt.Println("D")
	} else {
		fmt.Println("不及格")
	}
}

// if...if...else...else
// if嵌套语句
func ifNestedStatements() {
	/*
		var a int = 100
		var b int = 201

		if a == 100 {
			fmt.Println("a满足条件")
			if b == 200 {
				fmt.Println("b满足条件")
			}
		}
	*/

	// Demo -> 业务，验证密码正确
	var a, b int
	var pwd int = 20221020

	// 用户输入
	fmt.Println("请输入密码：")
	fmt.Scanln(&a)

	// 判断逻辑
	if a == pwd {
		fmt.Println("请再次输入密码")
		fmt.Scanln(&b)
		if a == b {
			fmt.Println("登录成功")
		} else {
			fmt.Println("登录失败")
		}
	} else {
		fmt.Println("密码错误")
	}

}
