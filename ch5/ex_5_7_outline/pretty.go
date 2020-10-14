package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth int

func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func start(n *html.Node) {
	switch {
	case n.Type == html.CommentNode:
	case n.Type == html.TextNode:
	case n.Type == html.ElementNode:
		startElement(n)
	}
}

func end(n *html.Node) {
	switch {
	case n.Type == html.CommentNode:
	case n.Type == html.TextNode:
	case n.Type == html.ElementNode:
		endElement(n)
	}
}

func startElement(n *html.Node) {
	var attrs []string
	for _, a := range n.Attr {
		attrs = append(attrs, fmt.Sprintf("%s='%s'", a.Key, a.Val))
	}

	attrsStr := strings.Join(attrs, " ")
	fmt.Printf("%*s<%s %s>\n", depth*2, "", n.Data, attrsStr)
	depth++
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	ForEachNode(doc, start, end)
}
