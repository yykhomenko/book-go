package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/yykhomenko/book-gopl/ch7/eval"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	if s.Scan(); s.Err() != nil {
		log.Fatalf("input: %v", s.Err())
	}

	e, err := eval.Parse(s.Text())
	if err != nil {
		log.Fatalf("parse: %v", err)
	}

	fmt.Fprintln(os.Stdout, e)
}
