// fetch golang.org | go run findlinks.go
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var attrs = map[string]string{
	"a":      "href",
	"img":    "src",
	"script": "src",
	"link":   "href",
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Fprintln(os.Stdout, link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && attrs[n.Data] != "" {
		for _, a := range n.Attr {
			if a.Key == attrs[n.Data] {
				links = append(links, a.Val)
			}
		}
	}

	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}

	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}
