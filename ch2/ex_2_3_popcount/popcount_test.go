// go test -bench .
package ex_2_3_popcount

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yykhomenko/book-gopl/ch2/popcount"
)

func TestPopCount(t *testing.T) {
	assert.Equal(t, 4, popcount.PopCount(15))
	assert.Equal(t, 4, PopCount(15))
}

// BenchmarkPopCount-8             1000000000               0.277 ns/op           0 B/op          0 allocs/op
func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(uint64(i))
	}
}

// BenchmarkPopCountCycle-8        70557334                16.4 ns/op             0 B/op          0 allocs/op
func BenchmarkPopCountCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}
