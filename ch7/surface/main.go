// go run main.go
// http://localhost:8000/plot?expr=sin(-x)*pow(1.5,-r)
// http://localhost:8000/plot?expr=pow(2,sin(y))*pow(2,sin(x))/12
// http://localhost:8000/plot?expr=sin(x*y/10)/10
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/yykhomenko/book-gopl/ch7/eval"
)

func main() {
	http.HandleFunc("/plot", plot)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "incorrect expression: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	printSVG(w, func(x, y float64) float64 {
		r := math.Hypot(x, y)
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
	})
}

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}

	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}

	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}

	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("unknown variable: %s", v)
		}
	}

	return expr, nil
}
