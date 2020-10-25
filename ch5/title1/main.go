package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if err := title(os.Args[1]); err != nil {
		log.Fatal(err)
	}
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s but not texl/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("analyze %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) bool {
		if n.Type == html.ElementNode &&
			n.Data == "title" &&
			n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
			return true
		}
		return false
	}

	forEachNode(doc, visitNode, nil)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil && pre(n) {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil && post(n) {
		return
	}
}
