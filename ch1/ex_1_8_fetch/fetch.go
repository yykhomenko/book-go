// Fetch makes HTTP GET queries, one by url.
// go build fetch.go && ./fetch example.com https://example1.com
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const prefix = "https://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v/n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: read: %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
