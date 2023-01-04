package main

import "fmt"

type dataset struct {
	name string
	age  int
}

func main() {
	ds := dataset{name: "", age: 1}
	fmt.Println(ds)
	change(&ds)
	fmt.Println(ds)
	fmt.Println("--------------")

	a := 1
	fmt.Println(a)
	changeTo(&a, 10)
	fmt.Println(a)
}

func changeTo(origin *int, to int) {
	*origin = to
}

func change(dataset *dataset) {
	dataset.name = "yahoo"
}
