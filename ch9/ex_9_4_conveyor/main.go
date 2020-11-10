package main

import (
	"fmt"
	"time"
)

func main() {
	in := make(chan int64)
	out := make(chan int64)

	go generateStage(in)
	go copyStage(out, in)
	printStage(out)
}

func generateStage(out chan<- int64) {
	defer close(out)
	out <- time.Now().UnixNano()
}

func copyStage(out chan<- int64, in <-chan int64) {
	for i := range in {
		out <- i
	}
}

func printStage(in <-chan int64) {
	start := <-in
	fmt.Printf("elapsed: %s", time.Duration(time.Now().UnixNano()-start))
}
