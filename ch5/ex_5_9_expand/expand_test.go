package ex_5_9_expand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {
	s := "Hello $user, you have $messCount unread messages"
	expected := "Hello Yurii, you have 2 unread messages"
	actual := expand(s, replacer)
	assert.Equal(t, expected, actual)
}

func replacer(arg string) string {
	switch arg {
	case "user":
		return "Yurii"
	case "messCount":
		return "2"
	default:
		return "<unknown>"
	}
}
