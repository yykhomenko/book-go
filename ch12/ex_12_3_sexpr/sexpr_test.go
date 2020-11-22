package sexpr

import (
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	test := struct {
		b bool
		f float64
		c complex128
		s []int
	}{
		true,
		1.25,
		1 + 2.5i,
		[]int{1, 2, 3},
	}

	b, err := Marshal(test)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(b))
}
