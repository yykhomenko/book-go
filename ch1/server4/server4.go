// Server4 is HTTP server that shows Lissajous figures.
// go run server4.go
// http :8000
package main

import (
	"log"
	"net/http"

	"github.com/yykhomenko/book-gopl/ch1/ex_1_6_lissajous/lissajous"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous.Lissajous(w)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
