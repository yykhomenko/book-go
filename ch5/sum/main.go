package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4))
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))

	var f func(...int)
	var g func([]int)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Pg. %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Println(os.Stderr)
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
