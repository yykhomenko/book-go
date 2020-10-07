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
	categories := make(map[string]int)
	invalid := 0
	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsSymbol(r) {
			categories["Symbol"]++
		}

		if unicode.IsLetter(r) {
			categories["Letter"]++
		}

		if unicode.IsMark(r) {
			categories["Mark"]++
		}

		if unicode.IsTitle(r) {
			categories["Title"]++
		}

		if unicode.IsSpace(r) {
			categories["Space"]++
		}

		if unicode.IsLower(r) {
			categories["Lower"]++
		}

		if unicode.IsUpper(r) {
			categories["Upper"]++
		}

		if unicode.IsPunct(r) {
			categories["Punct"]++
		}

		if unicode.IsNumber(r) {
			categories["Number"]++
		}

		if unicode.IsDigit(r) {
			categories["Digit"]++
		}

		if unicode.IsPrint(r) {
			categories["Print"]++
		}

		if unicode.IsControl(r) {
			categories["Control"]++
		}

		if unicode.IsControl(r) {
			categories["Control"]++
		}

		if unicode.IsGraphic(r) {
			categories["Graphic"]++
		}

		for name, table := range unicode.Properties {
			if unicode.In(r, table) {
				categories[name]++
			}
		}
	}

	fmt.Print("category\tcount\n")
	for c, n := range categories {
		fmt.Printf("%-30q\t%d\n", c, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d incorrect symbols UTF-8\n", invalid)
	}
}
