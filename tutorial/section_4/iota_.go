package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
	d
	e
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
