package sexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	test := struct {
		b bool
		f float64
		c complex128
		i interface{}
	}{
		true,
		1.25,
		1 + 2.5i,
		[]int{1, 2, 3},
	}

	expected := `((b t) (f 1.25) (c #C(1 2.5)) (i ("[]int" (1 2 3))))`

	b, err := Marshal(test)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, string(b))
}
