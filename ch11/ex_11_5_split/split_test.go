// go test -covermode=count -coverprofile=c.out
// go tool cover -html=c.out
package split_test

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s    string
		sep  string
		want int
	}{
		{"", ";", 1},
		{"a:b:c", ":", 3},
	}
	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q, %q) returns %d words, want %d", test.s, test.sep, got, test.want)
		}
	}
}
