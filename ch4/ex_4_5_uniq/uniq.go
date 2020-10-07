package ex_4_5_uniq

func uniq(strs []string) []string {
	var w int
	for _, s := range strs {
		if strs[w] != s {
			w++
			strs[w] = s
		}
	}
	return strs[:w+1]
}
