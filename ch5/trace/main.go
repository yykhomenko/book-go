package main

import (
	"log"
	"time"
)

func main() {
	bigSlowOperation()
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(2 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter in %s", msg)
	return func() {
		log.Printf("exit from %s (%s)", msg, time.Since(start))
	}
}
