// go run crawl.go -depth 2 https://golang.org
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"sync"

	"github.com/yykhomenko/book-gopl/ch5/links"
)

var depth = flag.Int("depth", math.MaxInt64, "search depth, unbounded by default")
var par = flag.Int("p", math.MaxInt64, "parallel factor, 20 by default")

func main() {
	flag.Parse()
	seen := make(map[string]bool)

	for _, link := range flag.Args() {
		dls(link, *depth, *par, seen, func(url string) []string {
			fmt.Println(url)
			urls, err := links.Extract(url)
			if err != nil {
				log.Print(err)
			}
			return urls
		})
	}
}

// dls is depth-limited search
func dls(s string, depth int, p int, seen map[string]bool, f func(string) []string) {
	if seen == nil {
		seen = make(map[string]bool)
	}

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	sm := make(chan struct{}, p)

	wg.Add(1)

	var dls_ func(s string, depth int)
	dls_ = func(s string, depth int) {
		defer wg.Done()
		if depth == 0 {
			return
		}

		sm <- struct{}{}
		strs := f(s)
		<-sm

		for _, link := range strs {
			mu.Lock()
			if seen[link] {
				mu.Unlock()
				continue
			}
			seen[link] = true
			mu.Unlock()
			wg.Add(1)
			go dls_(link, depth-1)
		}
	}

	dls_(s, depth)

	wg.Wait()
}
