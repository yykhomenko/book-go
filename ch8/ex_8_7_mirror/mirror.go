// go run crawl.go -depth 2 https://golang.org
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"

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
			log.Printf("downloaded %s", name)
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

func download(uri string) (filename string, err error) {
	resp, err := http.Get(uri)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("get %s: %v", uri, resp.Status)
	}

	link, err := url.Parse(uri)
	if err != nil {
		return "", err
	}

	filename = filepath.Join(link.Host, link.Path)
	if filepath.Ext(link.Path) == "" {
		filename = filepath.Join(link.Host, link.Path, "index.html")
	}

	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		return filename, err
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Print(err)
	}

	forAll(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for i, a := range n.Attr {
				if a.Key == "href" {
					if u, err := link.Parse(a.Val); err == nil {
						if u.Host == link.Host {
							u.Scheme = ""
							u.Host = ""
							u.User = nil
							u.Path = strings.TrimPrefix(u.Path+"index.html", "/")
							a.Val = u.String()
							n.Attr[i] = a
						}
					}
				}
			}
		}
	})

	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if err := html.Render(file, doc); err != nil {
		return "", nil
	}

	return filename, nil
}

func forAll(n *html.Node, f func(n *html.Node)) {
	f(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forAll(c, f)
	}
}
