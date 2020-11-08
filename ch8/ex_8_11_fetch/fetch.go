package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	result := make(chan []byte)
	cancel := make(chan struct{})
	for _, url := range os.Args[1:] {
		go fetch(url, result, cancel)
	}

	select {
	case b := <-result:
		close(cancel)
		os.Stdout.Write(b)
	}
}

func fetch(url string, result chan<- []byte, cancel <-chan struct{}) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	req.Cancel = cancel

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "get %s: %s\n", url, resp.Status)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	result <- b
}
