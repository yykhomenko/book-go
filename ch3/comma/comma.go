// "a/b/c.go" => "c"
// "c.d.go" => "c.d"
// "abc" => "abc"
package comma

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
