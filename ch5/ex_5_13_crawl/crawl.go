// Crawl downloads all pages from URL except foreign sites.
package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/yykhomenko/book-gopl/ch5/links"
)

func main() {
	urls := os.Args[1:]
	breadthFirst(crawlOnlyHosts(urls), urls)
}

func crawlOnlyHosts(urls []string) func(string) []string {
	var hosts []string
	for _, u := range urls {
		l, err := url.Parse(u)
		if err != nil {
			log.Fatalf("parse %s: %v/n", u, err)
		}
		hosts = append(hosts, l.Host)
	}

	return func(url string) []string {
		for _, host := range hosts {
			if strings.Contains(url, host) {
				fmt.Println(url)
				urls, err := links.Extract(url)
				if err != nil {
					log.Print(err)
				}
				return urls
			}
		}
		return nil
	}
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
