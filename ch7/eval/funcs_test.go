package eval

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	expr, err := Parse("x*y + z")
	if err != nil {
		log.Fatalf("parse: %v", err)
	}

	vars := Filter(expr, FilterVars)

	assert.Equal(t, Var("x"), vars[0])
	assert.Equal(t, Var("y"), vars[1])
	assert.Equal(t, Var("z"), vars[2])
}
