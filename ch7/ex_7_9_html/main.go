package main

import (
	"os"
	"sort"

	"github.com/yykhomenko/book-gopl/ch7/ex_7_9_html/sorting"
)

var tracks = []*sorting.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, sorting.Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, sorting.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, sorting.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, sorting.Length("4m24s")},
}

func main() {
	bc := sorting.NewByColumns(tracks, 3)
	bc.Select(sorting.LessYear)
	bc.Select(sorting.LessTitle)
	sort.Sort(bc)
	bc.PrintTracksHTML(os.Stdout)
}
