package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	x := IntSet{}
	x.Add(4)
	x.Add(2)
	assert.Equal(t, "{2 4}", x.String())
}

func TestHas(t *testing.T) {
	x := IntSet{}
	x.Add(4)
	x.Add(2)
	assert.True(t, x.Has(4))
	assert.True(t, !x.Has(3))
}
