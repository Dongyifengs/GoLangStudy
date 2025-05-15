package main

import "fmt"

func main() {
	//TIP 声明变量，并且这个变量要声明类型，还可以再次赋值。
	var name string = "MoYi"
	var age int
	name = "wzp"
	age = 19
	fmt.Println("Your Name is " + name + ".")
	fmt.Println("Your Age is", age)
}
