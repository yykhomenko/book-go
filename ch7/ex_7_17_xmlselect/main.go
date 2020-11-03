// curl -s https://www.w3.org/TR/2006/RECxmlll20060816 | go run main.go div div h2
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if err := FprintXMLTags(os.Stdout, os.Stdin, os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
		os.Exit(1)
	}
}

func FprintXMLTags(w io.Writer, r io.Reader, tags []string) error {
	dec := xml.NewDecoder(r)
	var stack [][]string
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			elem := []string{tok.Name.Local}
			for _, attr := range tok.Attr {
				elem = append(elem, attr.Name.Local)
			}
			stack = append(stack, elem) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, tags) {
				fmt.Fprintf(w, "%s: %s\n", strings.Join(flatten(stack), " "), tok)
			}
		}
	}
	return nil
}

func containsAll(x [][]string, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if matches(x[0], y[0]) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func matches(sElem []string, param string) bool {
	for _, e := range sElem {
		if e == param {
			return true
		}
	}
	return false
}

func flatten(stack [][]string) (res []string) {
	for _, elem := range stack {
		if len(elem) == 1 {
			res = append(res, elem[0])
		} else {
			res = append(res, elem[0]+"("+strings.Join(elem[1:], " ")+")")
		}
	}
	return
}
