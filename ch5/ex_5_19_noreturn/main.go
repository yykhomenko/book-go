package main

import "fmt"

func main() {
	fmt.Printf("return value: %d", noreturn())
}

func noreturn() (result int) {
	defer func() {
		if err := recover(); err != nil {
			result = 1
		}
	}()
	panic(0)
}
