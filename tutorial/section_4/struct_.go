package main

import "fmt"

type Person struct {
	first string
	last  string
}

type SecretAgent struct {
	Person
	ltk bool
}

func main() {
	sa1 := SecretAgent{
		Person: Person{},
		ltk:    false,
	}
	fmt.Println(sa1)

	p1 := Person{
		first: "Chris",
		last:  "Son",
	}
	fmt.Println(p1)

	aa(struct {
		Id   string
		name string
		age  int
	}{Id: "aa", name: "bb", age: 32})
}

func bb(a string) {

}

func aa(a struct {
	Id   string
	name string
	age  int
}) {
	fmt.Println(a)
}
