// Converter converts numeric argument to ˚C and ˚F.
// go run converter.go 32
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/yykhomenko/book-gopl/ch2/tempconv"
)

type Foot float64
type Meter float64
type Pound float64
type Kilogram float64

func (f Foot) String() string {
	return fmt.Sprintf("%gft", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}

func (p Pound) String() string {
	return fmt.Sprintf("%glb", p)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%gkg", k)
}

func main() {
	for _, arg := range os.Args[1:] {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(v)
		c := tempconv.Celsius(v)
		fmt.Fprintf(os.Stdout,
			"%s = %s, %s = %s\n",
			f, tempconv.FToC(f),
			c, tempconv.CToF(c),
		)

		foot := Foot(v)
		fmt.Fprintf(os.Stdout,
			"%s = %s, %s = %s\n",
			foot, FootToMeter(foot),
			foot, FootToMeter(foot),
		)
	}
}
