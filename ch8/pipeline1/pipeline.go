package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// generate
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// map to square
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// output
	for {
		fmt.Println(<-squares)
		time.Sleep(500 * time.Millisecond)
	}
}
