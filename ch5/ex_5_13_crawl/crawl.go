// Crawl downloads all pages from URL except foreign sites.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yykhomenko/book-gopl/ch5/links"
)

func main() {
	breadthFirst(crawl, os.Args[1:])
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func breadthFirst(f func(string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
