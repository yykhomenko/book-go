package ex_2_4_popcount

// PopCount returns number of bits set.
func PopCount(x uint64) int {
	var c uint64
	for i := 0; i < 64; i++ {
		c += (x >> i) & 1
	}
	return int(c)
}
