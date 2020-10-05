// go run anagram.go atomic cmiato //true
// go run anagram.go one eno //true
// go run anagram.go hello heloo //false
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, isAnagram(os.Args[1], os.Args[2]))
}

func isAnagram(a, b string) bool {
	aFreq := make(map[rune]int)
	for _, r := range a {
		aFreq[r]++
	}

	bFreq := make(map[rune]int)
	for _, r := range b {
		bFreq[r]++
	}

	for k, v := range aFreq {
		if bFreq[k] != v {
			return false
		}
	}

	for k, v := range bFreq {
		if aFreq[k] != v {
			return false
		}
	}

	return true
}
