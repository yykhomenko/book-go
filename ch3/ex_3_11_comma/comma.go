package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprintln(os.Stdout, comma("123"))
	fmt.Fprintln(os.Stdout, comma("1234"))
	fmt.Fprintln(os.Stdout, comma("12345"))
	fmt.Fprintln(os.Stdout, comma("123456"))
	fmt.Fprintln(os.Stdout, comma("-123456"))
	fmt.Fprintln(os.Stdout, comma("+123456"))
	fmt.Fprintln(os.Stdout, comma("1234.56"))
	fmt.Fprintln(os.Stdout, comma("-1234.56"))
	fmt.Fprintln(os.Stdout, comma("+1234.56"))
}

func comma(s string) string {
	var prefix string
	if '-' == s[0] || s[0] == '+' {
		prefix = string(s[0])
		s = s[1:]
	}

	var suffix string
	if i := strings.LastIndex(s, "."); i >= 0 {
		suffix = s[i:]
		s = s[:i]
	}

	return prefix + commaRec(s) + suffix
}

func commaRec(s string) string {
	const groupSize = 3
	n := len(s)
	if n <= groupSize {
		return s
	} else {
		return commaRec(s[:n-groupSize]) + "," + s[n-groupSize:]
	}
}
