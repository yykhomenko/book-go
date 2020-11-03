package xmltree

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXMLTree(t *testing.T) {
	input := `<doc><x unit="m">1</x><y unit="mm" id="height">2000</y></doc>`
	expected := "doc\n  x(unit=\"m\"): \"1\"\n  y(unit=\"mm\" id=\"height\"): \"2000\""
	tree, err := XMLTree(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expected, tree.(*Element).String())
}
