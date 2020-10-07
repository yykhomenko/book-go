package ex_4_6_onespace

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnespace(t *testing.T) {
	a := []byte("o  oone")
	b := []byte("o oone")
	c := onespace(a)
	fmt.Println(string(c))
	assert.Equal(t, b, c)
}
