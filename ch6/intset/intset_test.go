package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	s := IntSet{}
	s.Add(4)
	s.Add(2)
	assert.Equal(t, "{2 4}", s.String())
}

func TestHas(t *testing.T) {
	s := IntSet{}
	s.Add(4)
	s.Add(2)
	assert.True(t, s.Has(4))
	assert.True(t, !s.Has(3))
}
