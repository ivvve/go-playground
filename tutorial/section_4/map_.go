package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 32,
		"Chris": 27,
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println("Key:", k, "/ Value:", v)
	}

	fmt.Println("------------------------------")

	m["James2"] = 13
	v, ok := m["James2"]
	fmt.Println(v, ok)

	if v, ok := m["James"]; ok {
		fmt.Println(v)
	}

	fmt.Println("------")
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("------")

	delete(m, "James2")
	v, ok = m["James2"]
	fmt.Println(v, ok)
}
