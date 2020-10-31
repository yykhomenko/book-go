package counters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCounter(t *testing.T) {
	const input = "Hello world 2020 5678"
	var c WordCounter
	c.Write([]byte(input))
	c.Write([]byte(input))
	assert.Equal(t, 8, int(c))
}

func TestLineCounter(t *testing.T) {
	const input = "Hello\nworld\n2020\n5678"
	var c WordCounter
	c.Write([]byte(input))
	c.Write([]byte(input))
	assert.Equal(t, 8, int(c))
}
