package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	tests := []struct {
		v    interface{}
		want string
	}{
		{struct{}{}, "{}"},
		{[]struct{}{}, "[]"},
		{[]struct{ a int }{{1}, {2}}, `[{"a":1},{"a":2}]`},
	}
	for _, test := range tests {
		got, err := Marshal(test.v)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, test.want, string(got))
	}
}
