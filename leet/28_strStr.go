package leet

func strStr(haystack string, needle string) int {
	h := len(haystack)
	n := len(needle)
	if h < n {
		return -1
	}
	end := h - n + 1
	for i := 0; i < end; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}

	return -1
}
