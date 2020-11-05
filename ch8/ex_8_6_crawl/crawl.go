// go run crawl.go -depth 2 https://golang.org
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/yykhomenko/book-gopl/ch5/links"
)

var depth = flag.Int("depth", 0, "search depth, 0 is unbounded")

func main() {
	flag.Parse()

	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() { worklist <- flag.Args() }()

	crawlTo(worklist, unseenLinks)
	filterUnseen(unseenLinks, worklist)
}

func crawlTo(worklist chan<- []string, unseenLinks <-chan string) {
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				fmt.Println(link)
				foundLinks, err := links.Extract(link)
				if err != nil {
					log.Print(err)
				}
				go func() { worklist <- foundLinks }()
			}
		}()
	}
}

func filterUnseen(unseenLinks chan<- string, worklist <-chan []string) {
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
