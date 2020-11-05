// go run crawl.go https://golang.org
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yykhomenko/book-gopl/ch5/links"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)

	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens

	if err != nil {
		log.Print(err)
	}

	return list
}

func main() {
	n := 0
	lists := make(chan []string)

	n++
	go func() { lists <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-lists
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					lists <- crawl(link)
				}(link)
			}
		}
	}
}
