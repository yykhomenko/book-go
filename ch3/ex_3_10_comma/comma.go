package main

import (
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	const groupSize = 3
	n := len(s)
	if n <= groupSize {
		return s
	}
	var buf bytes.Buffer
	offset := n % groupSize
	buf.WriteString(s[:offset])
	for i := offset; i < n; i += groupSize {
		if i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[i : i+groupSize])
	}
	return buf.String()
}

func main() {
	fmt.Fprintln(os.Stdout, comma("123"))
	fmt.Fprintln(os.Stdout, comma("1234"))
	fmt.Fprintln(os.Stdout, comma("12345"))
	fmt.Fprintln(os.Stdout, comma("123456"))
}
