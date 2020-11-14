// cat charcount.go | go run charcount.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	printResults(charCount(in))
}

func charCount(in *bufio.Reader) (counts map[rune]int, utfLen [5]int, invalid int) {
	counts = make(map[rune]int)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utfLen[n]++
	}
	return
}

func printResults(counts map[rune]int, utfLen [5]int, invalid int) {
	fmt.Print("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utfLen {
		fmt.Printf("%d\t%d\n", i, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d incorrect symbols UTF-8\n", invalid)
	}
}
