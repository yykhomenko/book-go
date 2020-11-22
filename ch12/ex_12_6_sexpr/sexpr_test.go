package sexpr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yykhomenko/book-gopl/ch6/geometry"
)

func TestMarshal(t *testing.T) {
	test := struct {
		b     bool
		f     float64
		c     complex128
		i     interface{}
		m     map[string]string
		point geometry.Point
	}{
		true,
		0.0,
		0 + 0.0i,
		[]int{1, 2, 3},
		make(map[string]string),
		geometry.Point{1, 2},
	}

	test.m["a"] = "b"

	expected := `((b t)
 (i ("[]int" (1 2 3)))
 (m (("a" "b")))
 (point ((X 1)
         (Y 2))))`

	b, err := Marshal(test)
	if err != nil {
		t.Error(err)
	}

	r := string(b)
	fmt.Printf("%+v\n", test)
	fmt.Println(r)
	assert.Equal(t, expected, r)
}
