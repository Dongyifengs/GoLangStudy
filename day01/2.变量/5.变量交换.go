package main

import "fmt"

func main() {
	a := 100
	b := 200
	fmt.Printf("a:%d,b:%d\n", a, b)
	b, a = a, b
	fmt.Printf("a:%d,b:%d\n", a, b)
}
