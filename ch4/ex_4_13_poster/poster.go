// go build -o poster && ./poster predator
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const usage = `poster MOVIE_TITLE`
const UrlFormat = "https://www.omdbapi.com/?apikey=%x&t=%s"

type Movie struct {
	Title  string
	Year   string
	Poster string
}

func (m Movie) Filename() string {
	i := strings.LastIndex(m.Poster, ".")
	return fmt.Sprintf("%s_(%s)%s", m.Title, m.Year, m.Poster[i:])
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal(usage)
	}

	m, err := GetMovie(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	if err := WritePoster(m); err != nil {
		log.Fatal(err)
	}
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

func WritePoster(m *Movie) error {
	resp, err := http.Get(m.Poster)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unable to get poster: %v\n", resp.Status)
	}

	f, err := os.Create(m.Filename())
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	r := bufio.NewReader(resp.Body)

	if _, err = w.ReadFrom(r); err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return err
	}

	return nil
}
