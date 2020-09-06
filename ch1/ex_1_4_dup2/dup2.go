package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	var dupFileNames []string

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			if countLines(f, counts) {
				dupFileNames = append(dupFileNames, arg)
			}
			f.Close()
		}
	}

	for _, filename := range dupFileNames {
		fmt.Println("duplicates founded in file: ", filename)
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) bool {
	input := bufio.NewScanner(f)
	foundDup := false
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if counts[line] > 1 {
			foundDup = true
		}
	}

	return foundDup
}
