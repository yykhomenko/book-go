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
<style><a href="https://example.org/link1"</a>1</style>
<script><a href="https://example.org/link2"</a>2</script>
</head>
<body>
<a href="https://example.org/link3">3</a>
<p>4</p>
<a href="https://example.org/link4">5</a>
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
	expected := []string{"3", "4", "5"}
	assert.Equal(t, expected, values)
}
