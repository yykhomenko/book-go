// go run crawl.go -depth 2 https://golang.org
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strings"

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
			urls := filterByPrefixes(crawl(url), flag.Args())
			return urls
		})
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	urls, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return urls
}

func filterByPrefixes(strs, prefixes []string) (out []string) {
	for _, s := range strs {
		for _, p := range prefixes {
			if strings.HasPrefix(s, p) {
				out = append(out, s)
			}
		}
	}
	return
}
