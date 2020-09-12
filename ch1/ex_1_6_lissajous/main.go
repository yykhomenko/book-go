// Lissajous generate animation GIF file from random Lissajous figure.
// go run Lissajous.go > example.gif
package main

import (
	"os"

	"github.com/yykhomenko/book-gopl/ch1/ex_1_6_lissajous/lissajous"
)

func main() {
	lissajous.Lissajous(os.Stdout)
}
