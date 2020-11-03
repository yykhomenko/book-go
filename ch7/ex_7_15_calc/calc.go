package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/yykhomenko/book-gopl/ch7/eval"
)

func main() {
	fmt.Println("Expression:")
	s := bufio.NewScanner(os.Stdin)
	if s.Scan(); s.Err() != nil {
		log.Fatalf("input: %v", s.Err())
	}

	expr, err := eval.Parse(s.Text())
	if err != nil {
		log.Fatalf("parse: %v", err)
	}

	vars := eval.Filter(expr, eval.FilterVars)
	env := eval.Env{}

	fmt.Println("Please input follow variables:")
	for i := 0; i < len(vars); {
		fmt.Printf("%s = ", vars[i])
		if s.Scan(); s.Err() != nil {
			fmt.Fprintf(os.Stderr, "input: %v\n", err)
			continue
		}
		v, err := strconv.ParseFloat(s.Text(), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse: %v\n", err)
			continue
		}
		env[eval.Var(vars[i].String())] = v
		i++
	}

	fmt.Printf("%s = %g", expr, expr.Eval(env))
}
