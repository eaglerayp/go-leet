package ferrum

// func substringCalculator(s string) int64 {
// 	// Write your code here
// 	dup := map[string]bool{}
// 	dup[s] = true
// 	n := len(s)
// 	if n <= 1 {
// 		return int64(1)
// 	}
// 	recSubstring(s[0:n-1], dup)
// 	recSubstring(s[1:n], dup)
// 	return int64(len(dup))
// }

func recSubstring(s string, sm map[string]bool) {
	sm[s] = true
	n := len(s)
	if n <= 1 {
		return
	}
	recSubstring(s[0:n-1], sm)
	recSubstring(s[1:n], sm)
}

// caculate number of distince substring of string S
func substringCalculator(s string) int64 {
	// Write your code here
	n := len(s)
	var sum int64 = 0
	if n <= 1 {
		return int64(1)
	}
	for i := n; i > 0; i-- {
		dup := map[string]bool{}
		for j := 0; j < n-i+1; j++ {
			subS := s[j : i+j]
			dup[subS] = true
		}
		sum = sum + int64(len(dup))
	}
	return sum
}
