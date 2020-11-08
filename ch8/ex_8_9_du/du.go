// go run du.go -v ~ /
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "print intermediate results")

type sizeResp struct {
	root string
	size int64
}

type counter struct {
	nfiles, nbytes int64
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan sizeResp)
	wg := &sync.WaitGroup{}
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, root, wg, fileSizes)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	counters := make(map[string]*counter)
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			c, ok := counters[size.root]
			if !ok {
				c = &counter{}
				counters[size.root] = c
			}
			c.nfiles++
			c.nbytes += size.size
		case <-tick:
			printDiskUsage(counters)
		}
	}
	printDiskUsage(counters)
}

func walkDir(dir string, root string, wg *sync.WaitGroup, fileSizes chan<- sizeResp) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, root, wg, fileSizes)
		} else {
			fileSizes <- sizeResp{root, entry.Size()}
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(counters map[string]*counter) {
	for root, c := range counters {
		fmt.Printf("%s %d files %.1f GB\n", root, c.nfiles, float64(c.nbytes)/1e9)
	}
}
