package main

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCharCount(t *testing.T) {
	var tests = []struct {
		input       string
		wantCounts  map[rune]int
		wantUtfLen  [5]int
		wantInvalid int
	}{
		{"", map[rune]int{}, [5]int{}, 0},
		{"a", map[rune]int{'a': 1}, [5]int{1: 1}, 0},
		{"aa", map[rune]int{'a': 2}, [5]int{1: 2}, 0},
		{"ab", map[rune]int{'a': 1, 'b': 1}, [5]int{1: 2}, 0},
		{"été", map[rune]int{'t': 1, 'é': 2}, [5]int{1: 1, 2: 2}, 0},
	}

	for _, test := range tests {
		c, u, i := CharCount(strings.NewReader(test.input))
		if !cmp.Equal(test.wantCounts, c) {
			t.Errorf("CharCount(%q) got c = %v, wants %v \n", test.input, c, test.wantCounts)
		}
		if !cmp.Equal(test.wantUtfLen, u) {
			t.Errorf("CharCount(%q) got u = %v, wants %v \n", test.input, u, test.wantUtfLen)
		}
		if !cmp.Equal(test.wantInvalid, i) {
			t.Errorf("CharCount(%q) got i = %v, wants %v \n", test.input, i, test.wantInvalid)
		}
	}
}
