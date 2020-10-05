// go run anagram.go ana //true
// go run anagram.go xyyx //true
// go run anagram.go hello //false
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, isAnagram(os.Args[1]))
}

func isAnagram(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
