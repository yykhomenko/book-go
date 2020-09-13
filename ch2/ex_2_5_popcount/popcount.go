package ex_2_5_popcount

// PopCount returns number of bits set.
func PopCount(x uint64) int {
	var c int

	for x > 0 {
		x &= x - 1
		c++
	}

	return c
}
