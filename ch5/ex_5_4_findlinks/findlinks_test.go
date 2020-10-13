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
<a href="https://example.org/example"></a>
<img src="https://example.org/example.jpg"></a>
</body>
</html>
`)

func Test_visit(t *testing.T) {
	r := bytes.NewReader(HTMLText)
	doc, err := html.Parse(r)

	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	values := visit(nil, doc)
	expected := []string{
		"https://example.org/example.js",
		"example.css",
		"https://example.org/example",
		"https://example.org/example.jpg",
	}
	assert.Equal(t, expected, values)
}
