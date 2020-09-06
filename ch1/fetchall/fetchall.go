// Fetch makes parallel HTTP GET queries and shows elapsed time and body size.
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

	for range os.Args[1:] {
		fmt.Fprintln(os.Stdout, <-ch)
	}

	fmt.Fprintf(os.Stdout, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
