package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const usage = `poster MOVIE_TITLE`
const UrlFormat = "https://www.omdbapi.com/?apikey=%x&t=%s"

type Movie struct {
	Title  string
	Year   string
	Poster string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal(usage)
	}

	fmt.Println(GetMovie(os.Args[1]))
}

func GetMovie(title string) (*Movie, error) {
	url_ := fmt.Sprintf(UrlFormat, 0x253d414a, title)
	resp, err := http.Get(url_)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable to GET: %s\n", resp.Status)
	}

	movie := &Movie{}
	if err := json.NewDecoder(resp.Body).Decode(movie); err != nil {
		return nil, err
	}

	return movie, nil
}
