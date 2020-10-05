// "a/b/c.go" => "c"
// "c.d.go" => "c.d"
// "abc" => "abc"
package basename1

func basename(s string) string {
	for i := len(s); i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s); i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
