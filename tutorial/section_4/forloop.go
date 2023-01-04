package main

import "fmt"

func main() {
	switch {
	case true:
		d := "aa"
		fmt.Println(d)
	case false:
		d := "bb"
		fmt.Println(d)
	}
}
