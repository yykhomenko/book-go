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
		{true},
		{map[string]float64{"1": 1.2}},
		{map[int]int{1: 1}},
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
