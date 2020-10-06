package main

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		a, want [6]int
	}{
		{[6]int{0, 1, 2, 3, 4, 5}, [6]int{5, 4, 3, 2, 1, 0}},
	}
	for _, test := range tests {
		reverse(&test.a)
		if !reflect.DeepEqual(test.a, test.want) {
			t.Errorf("got %v, want %v", test.a, test.want)
		}
	}
}
