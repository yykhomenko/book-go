package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

var HTMLText = []byte(`
<a href="https://example.org/link"</a>
<a href="https://example.org/link"</a>
<a href="https://example.org/link"</a>
`)

func Test_visit(t *testing.T) {
	r := bytes.NewReader(HTMLText)
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	tags := visit(map[string]int{}, doc)
	assert.Equal(t, 3, tags["a"])
}
