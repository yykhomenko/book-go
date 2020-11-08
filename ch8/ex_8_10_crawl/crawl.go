// go run crawl.go -depth 2 https://golang.org
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"sync"

	"github.com/yykhomenko/book-gopl/ch5/links"
)

var (
	depth     = flag.Int("depth", math.MaxInt64, "search depth, empty is unbounded")
	semaphore = make(chan struct{}, 20)
	seen      = make(map[string]bool)
	seenMu    = &sync.Mutex{}
	done      = make(chan struct{})
)

func main() {
	flag.Parse()

	wg := &sync.WaitGroup{}
	for _, arg := range flag.Args() {
		wg.Add(1)
		crawlDeep(arg, *depth, wg)
	}

	go func() {
		wg.Wait()
		done <- struct{}{}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	select {
	case <-done:
		return
	case <-interrupt:
		close(done)
	}
}

func crawlDeep(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth == 0 {
		return
	}

	fmt.Println(url)

	semaphore <- struct{}{}
	foundLinks, err := links.ExtractWithCancel(url, done)
	<-semaphore
	if err != nil {
		log.Print(err)
	}

	for _, link := range foundLinks {
		seenMu.Lock()
		if seen[link] {
			seenMu.Unlock()
			continue
		}
		seen[link] = true
		seenMu.Unlock()
		wg.Add(1)
		go crawlDeep(link, depth-1, wg)
	}
}
