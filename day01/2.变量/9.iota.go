package main

import "fmt"

func main() {

	const (
		a = iota // a=0
		b
		c
		d = "hahaha"
		e
		f = 100
		g
		h = iota
		i
	)

	const (
		j = iota
		k
	)

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k)

}
