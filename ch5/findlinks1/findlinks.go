package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {

	f, err := os.Open("example.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	doc, err := html.Parse(f)
	// doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Fprintln(os.Stdout, link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		if n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = n.NextSibling {
		links = visit(links, c)
	}

	return links
}
