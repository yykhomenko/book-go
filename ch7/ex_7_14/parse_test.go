package eval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	const input = "-2 + <1, 4, -1, -3/4> * 2"
	const expected = -4.0
	expr, err := Parse(input)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expected, expr.Eval(Env{}))
}
