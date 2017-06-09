package algo

func longestPalindrome(s string) (result string) {
	if len(s) < 2 {
		return s
	}

	mL := 1
	n := len(s)
	result = s[:1]

	// traverse the string
	for i := 1; i < n; i++ {
		// challenge the longest case

		// version 1 but slower
		// for nowL := i + 1; nowL > mL; nowL-- {
		// 	if subS := s[i-nowL+1 : i+1]; isPalindrome(subS) {
		// 		mL = nowL
		// 		result = subS
		// 		continue
		// 	}
		// }

		// version two, each i only compute twice
		if isPalindromeIndex(s, i-mL, i) {
			// max length +1 case
			result = s[i-mL : i+1]
			mL = mL + 1
		} else if isPalindromeIndex(s, i-mL-1, i) {
			// max length +2 case
			result = s[i-mL-1 : i+1]
			mL = mL + 2
		}
	}
	return
}
