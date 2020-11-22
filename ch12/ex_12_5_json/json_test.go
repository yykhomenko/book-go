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
	}
	for _, test := range tests {
		got, err := Marshal(test.v)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, test.want, string(got))
	}
}
