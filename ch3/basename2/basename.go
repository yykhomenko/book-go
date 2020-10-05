// "a/b/c.go" => "c"
// "c.d.go" => "c.d"
// "abc" => "abc"
package basename1

import "strings"

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
