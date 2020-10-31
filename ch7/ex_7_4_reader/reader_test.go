package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestNewReader(t *testing.T) {
	data := "<title>Example</title>"
	doc, err := html.Parse(NewReader(data))
	assert.NoError(t, err)
	assert.Equal(t, "Example", doc.FirstChild.FirstChild.FirstChild.FirstChild.Data)
}
