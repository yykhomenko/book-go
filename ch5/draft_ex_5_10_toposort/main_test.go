package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_topoSort(t *testing.T) {
	expected := []string{
		"intro to programming",
		"discrete math",
		"data structures",
		"algorithms",
		"linear algebra",
		"calculus",
		"formal languages",
		"computer organization",
		"compilers",
		"databases",
		"operating systems",
		"networks",
		"programming languages",
	}
	actual := topoSort(prereqs)
	assert.Equal(t, expected, actual)
}
