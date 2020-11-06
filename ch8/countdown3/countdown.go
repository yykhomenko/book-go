package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Begin countdown. Press <Enter> to abort.")

	abort := make(chan struct{})
	launch := time.Tick(10 * time.Second)

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	select {
	case <-abort:
		fmt.Println("Launch aborted.")
		return
	case <-launch:
		fmt.Println("Launch!!!")
		// lanch()
	}
}
