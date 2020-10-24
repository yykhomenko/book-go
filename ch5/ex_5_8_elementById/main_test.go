package main

import (
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestElementByID(t *testing.T) {
	id := "title"
	input := `
<h1 id="test">example</h1>
<h2 id="` + id + `">example</h1>`

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	element := ElementByID(doc, id)
	assert.Equal(t, element.Data, "h2")
}
