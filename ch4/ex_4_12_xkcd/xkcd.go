// go build xkcd.go && ./xkcd init
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const UrlFormat = "https://xkcd.com/%d/info.0.json"

var usage = `xkcd init|search [(arg1 arg2 ...)]`

type Comics struct {
	Num        int
	Img        string
	Transcript string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal(usage)
	}
	cmd := os.Args[1]
	args := os.Args[2:]

	switch {
	case cmd == "init":
		initDB()
	case cmd == "search":
		search(args)
	}
}

func initDB() {
	f, err := os.Create("db.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for i := 1; i < 2; i++ {
		comics, err := GetComics(i)
		if err != nil {
			log.Fatal(err)
		}

		if err := json.NewEncoder(f).Encode(comics); err != nil {
			log.Fatal(err)
		}
	}
}

func search(args []string) {
	comics, err := GetDBComics()
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range comics {
		for _, arg := range args {

			if strings.Contains(c.Transcript, args[0]) {

			}
		}
	}
}

func GetDBComics() ([]Comics, error) {
	f, err := os.Open("db.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var comics []Comics
	if err := json.NewDecoder(f).Decode(&comics); err != nil {
		return nil, err
	}

	return comics, nil
}

func GetComics(n int) (*Comics, error) {
	url := fmt.Sprintf(UrlFormat, n)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable get data: %v\n", err)
	}

	comics := &Comics{}
	if err := json.NewDecoder(resp.Body).Decode(comics); err != nil {
		return nil, err
	}
	return comics, nil
}
