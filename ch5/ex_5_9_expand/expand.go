package ex_5_9_expand

import "strings"

func expand(s string, f func(string) string) string {
	return strings.ReplaceAll(s, "$foo", f("foo"))
}
