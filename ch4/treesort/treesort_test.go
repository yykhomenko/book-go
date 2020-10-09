package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	a := []int{5, 3, 4, 0, 2, 1}
	b := []int{0, 1, 2, 3, 4, 5}
	Sort(a)
	assert.Equal(t, b, a)
}
