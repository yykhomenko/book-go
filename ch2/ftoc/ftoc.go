// Ftoc converts two temperatures from ˚F to ˚C.
package main

import (
	"fmt"
	"os"
)

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Fprintf(os.Stdout, "%g˚F = %g˚C\n", freezingF, fToC(freezingF))
	fmt.Fprintf(os.Stdout, "%g˚F = %g˚C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
