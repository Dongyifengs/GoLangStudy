package main

import "fmt"

func main() {
	// bool的默认为false
	var b1 bool
	var b2 bool

	b1 = true
	b2 = false

	fmt.Println(b1, b2)
	fmt.Printf("%T,%t\n", b1, b1)
	fmt.Printf("%T,%t\n", b2, b2)
}
