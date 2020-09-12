package main

import (
	"fmt"
	"os"
)

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Fprintf(os.Stdout, "Temperature of boiling = %g˚F or %g˚C\n", f, c)
}
