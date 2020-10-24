package ex_5_9_expand

import (
	"regexp"
)

var r = regexp.MustCompile(`\$\w+`)

func expand(src string, f func(string) string) string {
	return r.ReplaceAllStringFunc(src, func(s string) string {
		return f(s[1:])
	})
}
