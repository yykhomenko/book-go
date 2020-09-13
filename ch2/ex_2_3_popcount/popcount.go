package ex_2_3_popcount

// pc[i] - number of single bits in i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns number of bits set.
func PopCount(x uint64) int {
	var c byte
	for i := 0; i < 8; i++ {
		c += pc[byte(x>>(8*i))]
	}
	return int(c)
}
