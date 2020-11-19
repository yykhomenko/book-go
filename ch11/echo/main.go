// Echo4 shows formatted args string
// go run main.go a b c
// go run main.go -s 2 a b c
// go run main.go -help
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	n   = flag.Bool("n", false, "skip new line symbol")
	sep = flag.String("s", " ", "separator")
)

var out io.Writer = os.Stdout

func main() {
	flag.Parse()
	if err := echo(!*n, *sep, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func echo(newline bool, sep string, args []string) error {
	fmt.Fprintf(out, "%s", strings.Join(args, sep))
	if newline {
		fmt.Fprintf(out, "%s", "\n")
	}
	return nil
}
