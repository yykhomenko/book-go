// go run sha256diff.go 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Input two arguments")
		os.Exit(1)
	}
	a := sha256.Sum256([]byte(os.Args[1]))
	b := sha256.Sum256([]byte(os.Args[2]))

	fmt.Fprintln(os.Stdout, diff(a, b))
}

func diff(a, b [32]byte) int {
	var c int
	for i := range a {
		c += byteDiff(a[i], b[i])
	}
	return c
}

func byteDiff(a, b byte) int {
	var c int
	for i := 0; i < 8; i++ {
		bitA := (a >> i) & 1
		bitB := (b >> i) & 1
		if bitA != bitB {
			c++
		}
	}
	return c
}
