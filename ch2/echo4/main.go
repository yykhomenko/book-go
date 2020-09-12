// Echo4 shows formatted args string
// go run main.go a b c
// go run main.go -s 2 a b c
// go run main.go -help
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var n = flag.Bool("n", false, "skip new line symbol")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Fprintf(os.Stdout, "%s", strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Fprintf(os.Stdout, "%s", "\n")
	}
}
