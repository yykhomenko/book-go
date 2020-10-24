package ex_5_9_expand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {
	s := "replace $foo and $foo"
	actual := expand(s, double)
	expected := "replace foofoo and foofoo"
	assert.Equal(t, expected, actual)
}

func double(arg string) string {
	return arg + arg
}
