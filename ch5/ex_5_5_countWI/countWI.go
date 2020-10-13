// go run countWI.go https://golang.org
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "count %s: %v\n", url, err)
			continue
		}
		fmt.Println(words, images)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	if n.Type == html.TextNode {
		data := strings.TrimSpace(n.Data)
		if len(data) != 0 {
			in := bufio.NewScanner(strings.NewReader(data))
			in.Split(bufio.ScanWords)
			for in.Scan() {
				words++
			}
		}
	}

	if n.FirstChild != nil {
		w, i := countWordsAndImages(n.FirstChild)
		words += w
		images += i
	}

	if n.NextSibling != nil {
		w, i := countWordsAndImages(n.NextSibling)
		words += w
		images += i
	}

	return
}
