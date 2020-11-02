// go run main.go
// http://localhost:8000/list
// http://localhost:8000/price?item=socks
// http://localhost:8000/update?item=hat&price=10
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := NewDatabase(map[string]dollars{"shoes": 50, "socks": 5})
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	sync.Mutex
	prices map[string]dollars
}

func NewDatabase(prices map[string]dollars) *database {
	return &database{prices: prices}
}

func (db *database) list(w http.ResponseWriter, r *http.Request) {
	db.Lock()
	defer db.Unlock()
	for item, price := range db.prices {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db *database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	db.Lock()
	defer db.Unlock()
	price, ok := db.prices[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item not found: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db *database) create(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, err := strconv.ParseFloat(r.URL.Query().Get("price"), 32)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "unable to parse price: %v\n", err)
		return
	}

	db.Lock()
	defer db.Unlock()

	if _, ok := db.prices[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item already present: %q\n", item)
		return
	}
	db.prices[item] = dollars(price)
}

func (db *database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, err := strconv.ParseFloat(r.URL.Query().Get("price"), 32)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "unable to parse price: %v\n", err)
		return
	}

	db.Lock()
	defer db.Unlock()

	if _, ok := db.prices[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item not found: %q\n", item)
		return
	}
	db.prices[item] = dollars(price)
}
