// go build -o xkcd main.go && ./xkcd init
// go build -o xkcd main.go && ./xkcd search go
package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(Usage)
	}
	cmd := os.Args[1]
	args := os.Args[2:]

	switch {
	case cmd == "init":
		InitDB()
	case cmd == "search":
		Search(args)
	default:
		log.Fatal(Usage)
	}
}
