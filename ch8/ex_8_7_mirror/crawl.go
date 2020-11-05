// go run crawl.go -depth 2 https://golang.org
package main

import (
	"flag"
	"fmt"
	"log"
	"math"

	"github.com/yykhomenko/book-gopl/ch5/links"
	"github.com/yykhomenko/book-gopl/ch8/search"
)

var depth = flag.Int("depth", math.MaxInt64, "search depth, unbounded by default")
var par = flag.Int("p", math.MaxInt64, "parallel factor, 20 by default")

func main() {
	flag.Parse()
	seen := make(map[string]bool)

	for _, link := range flag.Args() {
		search.DLS(link, *depth, *par, seen, func(url string) []string {
			fmt.Println(url)
			urls, err := links.Extract(url)
			if err != nil {
				log.Print(err)
			}
			return urls
		})
	}
}
