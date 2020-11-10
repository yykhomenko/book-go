// go 1.15.3, 2.9 GHz Quad-Core Intel Core i7
// messages per second: 2300119
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var counter uint64

func main() {
	pinc := make(chan string)
	pongc := make(chan string)
	go start(pinc)
	go stats()
	go pin(pinc, pongc)
	pong(pongc, pinc)
}

func start(pinc chan string) {
	pinc <- "pin"
}

func stats() {
	ticker := time.Tick(1 * time.Second)
	for {
		select {
		case <-ticker:
			fmt.Printf("messages per second: %d\n", counter)
			for !atomic.CompareAndSwapUint64(&counter, counter, 0) {
			}
		}
	}
}

func pin(out chan<- string, in <-chan string) {
	for range in {
		out <- "pin"
	}
}

func pong(out chan<- string, in <-chan string) {
	for range in {
		atomic.AddUint64(&counter, 1)
		out <- "pong"
	}
}
