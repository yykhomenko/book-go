// SHA256Diff counts different bits in two digest.
package ex_4_1_sha256

func diff(a, b [32]byte) int {
	var c int
	for i := range a {
		c += byteDiff(a[i], b[i])
	}
	return c
}

func byteDiff(a, b byte) int {
	var c int
	for i := 0; i < 8; i++ {
		bitA := (a >> i) & 1
		bitB := (b >> i) & 1
		if bitA != bitB {
			c++
		}
	}
	return c
}
