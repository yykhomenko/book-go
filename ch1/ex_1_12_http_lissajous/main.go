// Server4 is HTTP server that shows Lissajous figures.
// Supported configuration over HTTP GET parameters.
// go run main.go
// http://localhost:8000/?cycles=5&res=0.001&size=100&nFrames=64&delay=8
// http://localhost:8000/?cycles=8&res=0.001&size=400&nFrames=100&delay=2

package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/yykhomenko/book-gopl/ch1/ex_1_12_http_lissajous/lissajous"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		lissajous.Lissajous(w, newConfig(r))
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func newConfig(r *http.Request) lissajous.Config {
	conf := lissajous.Config{}

	if s, ok := r.Form["cycles"]; ok {
		if v, err := strconv.Atoi(s[0]); err == nil {
			conf.Cycles = v
		}
	}
	if s, ok := r.Form["res"]; ok {
		if v, err := strconv.ParseFloat(s[0], 64); err == nil {
			conf.Res = v
		}
	}

	if s, ok := r.Form["size"]; ok {
		if v, err := strconv.Atoi(s[0]); err == nil {
			conf.Size = v
		}
	}

	if s, ok := r.Form["nFrames"]; ok {
		if v, err := strconv.Atoi(s[0]); err == nil {
			conf.NFrames = v
		}
	}

	if s, ok := r.Form["delay"]; ok {
		if v, err := strconv.Atoi(s[0]); err == nil {
			conf.Delay = v
		}
	}

	return conf
}
