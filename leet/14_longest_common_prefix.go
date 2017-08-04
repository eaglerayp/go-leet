package leet

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n < 1 {
		return ""
	}
	max := strs[0]
	for i := 1; i < n; i++ {
		mm := len(max)

		if mm == 0 {
			return max
		}
		ns := len(strs[i])
		if ns < mm {
			mm = ns
			max = max[:ns]
		}
		for j := 0; j < mm; j++ {
			if strs[i][j] != max[j] {
				max = max[:j]
				break
			}
		}
	}
	return max
}
