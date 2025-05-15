package main

import "fmt"

var name string

func main() {
	name = "MoYi"
	fmt.Println(name)
	aaa()
}

func aaa() {
	name = "Wzp"
	fmt.Println(name)
}
