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
<html>
<head>
<script src="https://example.org/example.js"></script>
<link rel="stylesheet" type="text/css" href="example.css">
</head>
<body>
<a href="https://example.org/example">hello world</a>
<img src="https://example.org/example.jpg">hello world</a>
</body>
</html>
`)

func Test_countWordsAndImages(t *testing.T) {
	r := bytes.NewReader(HTMLText)
	doc, err := html.Parse(r)

	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	words, images := countWordsAndImages(doc)

	assert.Equal(t, 4, words)
	assert.Equal(t, 1, images)
}
