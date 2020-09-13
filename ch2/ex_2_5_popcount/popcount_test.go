// go test
// go test -bench . -benchmem
package ex_2_5_popcount

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yykhomenko/book-gopl/ch2/popcount"
)

func TestPopCount(t *testing.T) {
	assert.Equal(t, 3, PopCount(69))
}

// BenchmarkPopCount-8             1000000000               0.277 ns/op           0 B/op          0 allocs/op
func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(uint64(i))
	}
}

// BenchmarkPopCountClear-8        139278291                9.00 ns/op            0 B/op          0 allocs/op
func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}
