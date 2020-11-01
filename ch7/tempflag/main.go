// go run main.go
// go run main.go -temp 212°F
package main

import (
	"flag"
	"fmt"

	tempconv "github.com/yykhomenko/book-gopl/ch7/ex_7_6_tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
