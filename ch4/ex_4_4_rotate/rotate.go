package ex_4_4_rotate

func rotate(ints []int, n int) {
	first := make([]int, n)
	copy(first, ints[:n])
	copy(ints, ints[n:])
	copy(ints[len(ints)-n:], first)
}
