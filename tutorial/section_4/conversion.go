package main

import "fmt"

type Money int

func main() {
	var a Money = 1
	var b = int(a)
	fmt.Println(b)
}
