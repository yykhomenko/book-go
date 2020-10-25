package ex_5_15_min_max

func min(arg1 int, args ...int) int {
	min := arg1
	for _, arg := range args {
		if min > arg {
			min = arg
		}
	}
	return min
}

func max(arg1 int, args ...int) int {
	max := arg1
	for _, arg := range args {
		if max < arg {
			max = arg
		}
	}
	return max
}
