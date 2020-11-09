// go test
// go test -bench . -benchmem
package popcount

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopCount(t *testing.T) {
	assert.Equal(t, 4, PopCount(15))
}

var result int

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = PopCount(uint64(i))
	}
}
