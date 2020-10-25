// Crawl downloads all pages from URL except foreign sites.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/yykhomenko/book-gopl/ch5/links"
)

func main() {
	urls := os.Args[1:]
	breadthFirst(crawlOnlyHosts(urls), urls)
}

func breadthFirst(f func(string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawlOnlyHosts(urls []string) func(string) []string {
	var hosts []string
	for _, u := range urls {
		l, err := url.Parse(u)
		if err != nil {
			log.Fatalf("parse %s: %v/n", u, err)
		}
		hosts = append(hosts, l.Host)
	}

	return func(link string) []string {
		for _, host := range hosts {
			if strings.Contains(link, host) {
				fmt.Println(link)

				if err := download(link); err != nil {
					log.Print(err)
				}

				urls, err := links.Extract(link)
				if err != nil {
					log.Print(err)
				}

				return urls
			}
		}
		return nil
	}
}

func download(link string) error {
	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("get %s: %v", link, resp.Status)
	}

	u, err := url.Parse(link)
	if err != nil {
		return err
	}

	path := strings.Join(strings.Split(u.Host+u.Path, "/"), string(os.PathSeparator))
	filename := path + string(os.PathSeparator) + "index.html"

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	return err
}
