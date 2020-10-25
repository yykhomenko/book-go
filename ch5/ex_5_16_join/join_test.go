package ex_5_16_join

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	assert.Equal(t, "Hello world !!!", Join(" ", "Hello", "world", "!!!"))
}
