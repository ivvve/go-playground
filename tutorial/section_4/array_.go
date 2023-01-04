package main

import "fmt"

func main() {
	arr := [3]int{}
	fmt.Println(arr)
	println(len(arr))
	println(cap(arr))

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	println(len(arr))
	println(cap(arr))
}
