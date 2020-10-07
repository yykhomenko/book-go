package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	s := []byte("Räksmörgås")
	got := string(revUTF8(s))
	want := "sågrömskäR"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
