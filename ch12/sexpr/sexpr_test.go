package sexpr

import (
	"fmt"
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

func TestUnmarshal(t *testing.T) {
	expected := testStruct{
		true,
		0.0,
		0 + 0.0i,
		[]int{1, 2, 3},
		map[string]string{"a": "b"},
		geometry.Point{1, 2},
	}
	input := `((b true)
 (i ("[]int" (1 2 3)))
 (m (("a" "b")))
 (point ((X 1)
         (Y 2))))`

	var actual testStruct
	err := Unmarshal([]byte(input), &actual)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", actual)
	assert.Equal(t, expected, actual)
}
