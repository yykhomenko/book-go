// cat wordfreq.go | go run wordfreq.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		counts[in.Text()]++
	}

	fmt.Fprint(os.Stdout, "word\tfreq\n")
	for w, n := range counts {
		fmt.Fprintf(os.Stdout, "%q\t%d\n", w, n)
	}
}
