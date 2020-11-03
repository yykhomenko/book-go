package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintXMLTags(t *testing.T) {
	r := strings.NewReader(`<div id="page" class="wide">Example</div>
<div id="page">Example2</div>`)
	expected := "div(id class): Example\n"
	var buf bytes.Buffer
	tags := []string{"class"}

	if err := FprintXMLTags(&buf, r, tags); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, buf.String())
}
