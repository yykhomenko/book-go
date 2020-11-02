package sorting

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByColumns(t *testing.T) {
	// "Title", "Artist", "Album", "Year", "Length"
	var tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	var expected = []*Track{
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}

	cs := NewByColumns(tracks, 3)
	cs.Select(LessYear)
	cs.Select(LessTitle)

	sort.Sort(cs)
	cs.PrintTracks()

	assert.Equal(t, expected, tracks)
}
