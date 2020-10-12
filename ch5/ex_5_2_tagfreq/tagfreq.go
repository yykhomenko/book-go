// fetch golang.org | go run tagfreq.go
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	for tag, count := range visit(map[string]int{}, doc) {
		fmt.Fprintf(os.Stdout, "%q %d\n", tag, count)
	}
}

func visit(tags map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		tags[n.Data]++
	}

	if n.FirstChild != nil {
		tags = visit(tags, n.FirstChild)
	}

	if n.NextSibling != nil {
		tags = visit(tags, n.NextSibling)
	}

	return tags
}
