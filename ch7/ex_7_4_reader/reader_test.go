package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestNewReader(t *testing.T) {
	data := "<title>Example</title>"
	r := NewReader(data)
	_, err := html.Parse(*r)
	assert.NoError(t, err)
}
