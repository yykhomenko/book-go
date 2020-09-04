package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s string

	for k, v := range os.Args[1:] {
		s += strconv.Itoa(k+1) + " " + v + "\n"
	}

	fmt.Println(s)
}
