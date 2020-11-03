// http://localhost:8000/calc?expr=x*y*sin(x)&x=2&y=3
package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/yykhomenko/book-gopl/ch7/eval"
)

func main() {
	http.HandleFunc("/calc", func(w http.ResponseWriter, r *http.Request) {
		exp := r.URL.Query().Get("expr")
		if exp == "" {
			http.Error(w, "param expr is empty", http.StatusBadRequest)
			return
		}

		expr, err := eval.Parse(exp)
		if err != nil {
			http.Error(w, fmt.Sprintf("parse expr %q: %v", exp, err), http.StatusBadRequest)
			return
		}

		env := eval.Env{}
		vars := eval.Filter(expr, eval.FilterVars)
		for _, v := range vars {
			value, err := strconv.ParseFloat(r.URL.Query().Get(v.String()), 64)
			if err != nil {
				http.Error(w, fmt.Sprintf("parse param %q: %v", v, err), http.StatusBadRequest)
				return
			}
			env[v.(eval.Var)] = value
		}

		fmt.Fprintln(w, expr.Eval(env))
	})

	http.ListenAndServe("localhost:8000", nil)
}
