// Cf converts numeric argument to ˚C and ˚F.
// go run cf.go 32
// go run cf.go 212
// go run cf.go -40
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/yykhomenko/book-gopl/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		fmt.Fprintf(os.Stdout,
			"%s = %s, %s = %s\n",
			f, tempconv.FToC(f),
			c, tempconv.CToF(c),
		)
	}
}
