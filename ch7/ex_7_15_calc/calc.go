package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	eval "github.com/yykhomenko/book-gopl/ch7/ex_7_13"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	if sc.Err() != nil {
		log.Fatalf("input: %v", sc.Err())
	}
	e, err := eval.Parse(sc.Text())
	if err != nil {
		log.Fatalf("parse: %v", err)
	}

	fmt.Fprintln(os.Stdout, e)
}
