package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yykhomenko/book-gopl/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	lists := make(chan []string)

	go func() { lists <- os.Args[1:] }()

	seen := make(map[string]bool)
	for list := range lists {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					lists <- crawl(link)
				}(link)
			}
		}
	}
}
