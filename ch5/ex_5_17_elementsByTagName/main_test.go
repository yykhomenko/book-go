package main

import (
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestElementByID(t *testing.T) {
	input := `
<h1>example</h1>
<h2>example</h2>
<h3>example</h3>`

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	elements := ElementsByTagName(doc, "h2", "h3")
	assert.Equal(t, 2, len(elements))
}
