// Htmldata prints value of text tags(except style and script section) to stdout.
// fetch golang.org | go run htmldata.go
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s\n", visit(nil, doc))
}

func visit(values []string, n *html.Node) []string {
	if n.Type == html.ElementNode &&
		(n.Data == "style" || n.Data == "script") {
		return values
	}

	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if len(text) > 0 {
			values = append(values, text)
		}
	}

	if n.FirstChild != nil {
		values = visit(values, n.FirstChild)
	}

	if n.NextSibling != nil {
		values = visit(values, n.NextSibling)
	}

	return values
}
