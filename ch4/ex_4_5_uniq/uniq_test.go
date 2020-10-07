package ex_4_5_uniq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	a := []string{"", "one", "two", "two", "three", "", "three", "three"}
	b := []string{"", "one", "two", "three", "", "three"}
	c := uniq(a)
	assert.Equal(t, b, c)
}
