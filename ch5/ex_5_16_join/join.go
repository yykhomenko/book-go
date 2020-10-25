package ex_5_16_join

import "strings"

func Join(sep string, elems ...string) string {
	return strings.Join(elems, sep)
}
