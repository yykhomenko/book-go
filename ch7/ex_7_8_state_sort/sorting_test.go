package sorting

import (
	"sort"
	"testing"
)

func TestStateSorting(t *testing.T) {
	var tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	ss := stateSort{tracks, nil}
	ss.By("Artist")
	sort.Sort(ss)
	ss.PrintTracks()
}
