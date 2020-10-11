package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	Usage     = `xkcd init|search [(arg1 arg2 ...)]`
	UrlCount  = "https://xkcd.com/info.0.json"
	UrlFormat = "https://xkcd.com/%d/info.0.json"
)

type Comics struct {
	Num        int
	Img        string
	Transcript string
}

func InitDB() {
	comicsCount := comicsCount()

	var result []Comics
	for i := 1; i <= comicsCount; i++ {
		cs, err := GetComics(i)
		if err != nil {
			log.Printf("%v\n", err)
			continue
		}
		result = append(result, *cs)
		fmt.Printf("\r%3.2f%% (%d of %d)... ",
			100*float64(i)/float64(comicsCount),
			i,
			comicsCount)
	}

	f, err := os.Create("db.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	if err := enc.Encode(result); err != nil {
		log.Fatal(err)
	}
}

func comicsCount() int {
	resp, err := http.Get(UrlCount)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("unable to get: %v\n", resp.Status)
	}

	comics := &Comics{}
	if err := json.NewDecoder(resp.Body).Decode(comics); err != nil {
		log.Fatal(err)
	}
	return comics.Num
}

func Search(args []string) {
	comics, err := GetDBComics()
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range comics {
		for _, arg := range args {
			if strings.Contains(c.Transcript, arg) {
				fmt.Printf("\n#%d\nlink: %s\ntxt: %s\n", c.Num, c.Img, c.Transcript)
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
		fmt.Println(comics)
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
		return nil, fmt.Errorf("unable get data for %d comics: %v", n, resp.Status)
	}

	comics := &Comics{}
	if err := json.NewDecoder(resp.Body).Decode(comics); err != nil {
		return nil, err
	}
	return comics, nil
}
