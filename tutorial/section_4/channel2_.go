package main

import "fmt"

func main() {
	c := make(chan int)

	go foo(c)
	bar(c)
}

func foo(c chan<- int) {
	c <- 42
	close(c) // for문으로 값을 가져오는 경우 channel을 닫아야한다
}

func bar(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
