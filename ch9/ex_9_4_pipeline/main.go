// go 1.15.3, macos
// 1e6 stages = 2.56GB mem, travel 400ms, 400ns per stage
// 2e6 stages = 5.12GB mem, travel 800ms, 400ns per stage
// 3e6 stages = 7.72GB mem, travel 1.14s, 380ns per stage
// 4e6 stages = 10.24GB mem, travel 9.4s, 2.3us per stage
package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"
)

const doStages = int64(4000000)

func main() {
	in := make(chan int64)
	var out chan int64
	start := time.Now()

	var tempIn = in
	for i := int64(0); i < doStages; i++ {
		out = make(chan int64)
		go doStage(out, tempIn)
		tempIn = out
	}
	fmt.Printf("builded %d stages pipeline in %s\n", doStages, time.Now().Sub(start))

	go scan(in)
	printer(out)
}

func scan(in chan int64) {
	func() {
		defer close(in)
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			in <- time.Now().UnixNano()
		}
	}()
}

func printer(in <-chan int64) {
	for start := range in {
		duration := time.Now().UnixNano() - start
		fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
		fmt.Printf("travel time: %s\n", time.Duration(duration))
		fmt.Printf("time per stage: %s\n", time.Duration(duration/doStages))
	}
	fmt.Printf("\ngoroutines on exit: %d\n", runtime.NumGoroutine())
}

func doStage(out chan<- int64, in <-chan int64) {
	defer close(out)
	for i := range in {
		out <- i
	}
}
