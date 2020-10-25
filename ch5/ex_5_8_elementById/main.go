// curl -s https://www.google.com | go run main.go gbv
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: curl -s example.com | elementById ID")
	}

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "HTML parse: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%+v", ElementByID(doc, os.Args[1]))
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, isAttrPresent(id), nil)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) (result *html.Node) {
	if pre != nil && pre(n) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result = forEachNode(c, pre, post)
		if result != nil {
			return
		}
	}

	if post != nil && post(n) {
		return n
	}

	return nil
}

func isAttrPresent(id string) func(n *html.Node) bool {
	return func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == id {
					return true
				}
			}
		}
		return false
	}
}
