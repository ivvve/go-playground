package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		i := 1
		for {
			time.Sleep(time.Millisecond * 1_200)
			c1 <- 1
			i++

			if i == 3 {
				close(c1)
				break
			}
		}
	}()

	go func() {
		i := 1
		for {
			time.Sleep(time.Millisecond * 1_500)
			c2 <- 2
			i++

			if i == 3 {
				close(c2)
				break
			}
		}
	}()

	for {
		//v1 := <- c1
		//fmt.Println(v1)

		select {
		case v1 := <-c1:
			fmt.Println("v1 -", v1)
		case v2 := <-c2:
			fmt.Println("v2 -", v2)
		}
	}
}
