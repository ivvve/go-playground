package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		c1 <- getA()
	}()
	go func() {
		c2 <- getB()
	}()

	fmt.Println(<-c1)
	fmt.Println(<-c2)
}

func getA() string {
	fmt.Println("a start")
	time.Sleep(time.Second * 1)
	fmt.Println("a end")
	return "a"
}

func getB() string {
	fmt.Println("b start")
	time.Sleep(time.Second * 2)
	fmt.Println("b end")
	return "b"
}
