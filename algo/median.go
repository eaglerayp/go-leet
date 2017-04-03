package algo

func isPalindrome(s string) bool {
	n := len(s)
	half := n / 2
	for i := 0; i < half; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) (result string) {
	if len(s) < 1 {
		return s
	}

	mL := 1
	n := len(s)
	result = s[:1]

	// traverse the string
	for i := 1; i < n; i++ {
		// challenge the longest case
		for nowL := i + 1; nowL > mL; nowL-- {
			if subS := s[i-nowL+1 : i+1]; isPalindrome(subS) {
				mL = nowL
				result = subS
				continue
			}
		}
	}
	// fmt.Println("S:", s)
	// fmt.Println("result:", result)
	return
}
