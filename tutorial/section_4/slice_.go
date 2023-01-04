package main

import "fmt"

func main() {
	slice := []int{}
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println()

	slice = append(slice, 1)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println()

	slice = append(slice, 2)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println()

	slice = append(slice, 3)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println()

	slice = append(slice, slice...)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println()

	fmt.Println(slice[:1])
	fmt.Println()

	s := make([]int, 1, 12)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println()
}
