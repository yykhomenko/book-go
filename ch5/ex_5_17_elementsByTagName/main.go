// curl -s https://www.google.com | go run main.go a img
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: curl -s example.com | ./elementsByTagName name1 name2 ...")
	}

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	for _, node := range ElementsByTagName(doc, os.Args[1:]...) {
		fmt.Fprintf(os.Stdout, "%+v\n", node)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	return forEachNode(doc, isTagPresent(name...))
}

func forEachNode(n *html.Node, pre func(n *html.Node) bool) (results []*html.Node) {
	if pre != nil && pre(n) {
		results = append(results, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		results = append(results, forEachNode(c, pre)...)
	}

	return results
}

func isTagPresent(names ...string) func(n *html.Node) bool {
	return func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, name := range names {
				if n.Data == name {
					return true
				}
			}
		}
		return false
	}
}
