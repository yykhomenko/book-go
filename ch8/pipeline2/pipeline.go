package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		defer close(naturals)
		for x := 0; x < 100; x++ {
			naturals <- x
		}
	}()

	go func() {
		defer close(squares)
		for {
			for x := range naturals {
				squares <- x * x
			}
		}
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
