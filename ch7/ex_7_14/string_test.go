package eval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallString(t *testing.T) {
	const input = "sin(x)+-25++30+sqrt(2)*2/10-3.14e0/pow(y/2,2)+<4, 1>"
	const expected = "sin(x) + -25 + +30 + sqrt(2)*2/10 - 3.14/pow(y/2, 2) + <4, 1>"
	expr, err := Parse(input)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expected, expr.String())
}
