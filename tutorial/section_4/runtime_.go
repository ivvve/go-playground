package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("ARCH", runtime.GOARCH)
	fmt.Println("# of CPU", runtime.NumCPU())
	go func() {}()
	go func() {}()
	go func() {}()
	fmt.Println("# of Goroutine", runtime.NumGoroutine())
	fmt.Println()
}
