// Fetch makes parallel HTTP GET queries and dump to file elapsed time and body size.
// go run fetchall.go https://example.com https://example1.com
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	file, err := os.Create("dump.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "create file: %v/n", err)
		os.Exit(1)
	}
	defer file.Close()

	for range os.Args[1:] {
		fmt.Fprintln(file, <-ch)
	}

	fmt.Fprintf(file, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	n, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}

	t := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", t, n, url)
}
