package myjson

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	tests := []struct {
		v interface{}
	}{
		{struct{}{}},
		{[]struct{}{}},
		// {[]struct{ a int }{{-1}, {2}}},
		// {[]struct{ a, b float64 }{{-1.15, 2.5}}},
		// {struct{ a, b float64 }{-1.15, 2.5}},
		{map[string]float64{"1": 1.2}},
	}
	for _, test := range tests {
		expected, err := json.Marshal(test.v)
		if err != nil {
			t.Error(err)
		}
		actual, err := Marshal(test.v)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, string(expected), string(actual))
	}
}
