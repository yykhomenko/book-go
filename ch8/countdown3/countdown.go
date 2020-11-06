package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Begin countdown. Press <Enter> to abort.")

	abort := make(chan struct{})
	tick := time.Tick(1 * time.Second)

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-abort:
			fmt.Println("Launch aborted.")
			return
		case <-tick:
			// NOP
		}
	}
}
