package counters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCounter(t *testing.T) {
	var c WordCounter
	const input = "Hello world 2020 5678"
	c.Write([]byte(input))
	c.Write([]byte(input))
	assert.Equal(t, 8, int(c))
}
