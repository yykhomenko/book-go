package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeString(t *testing.T) {
	var tr *tree
	Add(tr, 4)
	Add(tr, 4)
	Add(tr, 3)
	Add(tr, 5)
	Add(tr, 2)
	assert.Equal(t, "{1 2 3 4 5}", tr.String())
}
