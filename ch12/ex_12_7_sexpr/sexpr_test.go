package sexpr

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yykhomenko/book-gopl/ch6/geometry"
)

type testStruct struct {
	b     bool
	f     float64
	c     complex128
	i     interface{}
	m     map[string]string
	point geometry.Point
}

func TestMarshal(t *testing.T) {
	expected := &testStruct{
		true,
		1.0,
		1 + 0.25i,
		[]int{1, 2, 3},
		map[string]string{"a": "b"},
		geometry.Point{1, 2},
	}

	input := `((b t)
 (f 1)
 (c #C(1 0.25))
 (i ("[]int" (1 2 3)))
 (m (("a" "b")))
 (point ((X 1)
         (Y 2))))`

	actual := &testStruct{}
	d := json.NewDecoder(strings.NewReader(input))
	if err := d.Decode(actual); err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, actual)
}
