// go run crawl.go -depth 2 https://golang.org
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/yykhomenko/book-gopl/ch5/links"
	"github.com/yykhomenko/book-gopl/ch8/search"
)

var depth = flag.Int("depth", math.MaxInt64, "search depth, unbounded by default")
var par = flag.Int("p", math.MaxInt64, "parallel factor, 20 by default")

func main() {
	flag.Parse()
	seen := make(map[string]bool)

	for _, link := range flag.Args() {
		search.DLS(link, *depth, *par, seen, func(url string) []string {
			name, err := download(url)
			if err != nil {
				log.Printf("unable to download %s: %v", url, err)
			}
			log.Printf("save %s", name)
			return filterByPrefixes(crawl(url), flag.Args())
		})
	}
}

func filterByPrefixes(strs, prefixes []string) (out []string) {
	for _, s := range strs {
		for _, p := range prefixes {
			if strings.HasPrefix(s, p) {
				out = append(out, s)
			}
		}
	}
	return
}

func crawl(url string) []string {
	urls, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return urls
}

func download(link string) (filename string, err error) {
	resp, err := http.Get(link)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("get %s: %v", link, resp.Status)
	}

	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	path := strings.Join(strings.Split(u.Host+u.Path, "/"), string(os.PathSeparator))
	filename = path + string(os.PathSeparator) + "index.html"

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return filename, err
	}

	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	return filename, nil
}
