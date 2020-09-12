// Converter converts numeric argument to ˚C, ˚F, ft, m, lb, kg.
// go run converter.go 32 100
// echo "32 100" | go run converter.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/yykhomenko/book-gopl/ch2/tempconv"
)

type Foot float64
type Meter float64
type Pound float64
type Kilogram float64

const (
	MetersInFoot     = 0.3048
	KilogramsInPound = 0.45359237
)

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

func FootToMeter(f Foot) Meter {
	return Meter(MetersInFoot * f)
}

func MeterToFoot(m Meter) Foot {
	return Foot(m / MetersInFoot)
}

func PoundToKilogram(p Pound) Kilogram {
	return Kilogram(KilogramsInPound * p)
}

func KilogramToPound(k Kilogram) Pound {
	return Pound(k / KilogramsInPound)
}

func main() {
	var args []string

	if len(os.Args) == 1 {
		s, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "converver: %v\n", err)
			os.Exit(1)
		}

		text := strings.Replace(s, "\n", "", -1)
		args = strings.Split(text, " ")
	} else {
		args = os.Args[1:]
	}

	for _, arg := range args {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "converver: %v\n", err)
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
		meter := Meter(v)
		fmt.Fprintf(os.Stdout,
			"%s = %s, %s = %s\n",
			foot, FootToMeter(foot),
			meter, MeterToFoot(meter),
		)

		pound := Pound(v)
		kilogram := Kilogram(v)
		fmt.Fprintf(os.Stdout,
			"%s = %s, %s = %s\n",
			pound, PoundToKilogram(pound),
			kilogram, KilogramToPound(kilogram),
		)
	}
}
