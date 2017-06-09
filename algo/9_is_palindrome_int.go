package algo

func isPalindromeInteger(x int) bool {
	// only compare half of digit
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	result := 0
	for x > result {
		result = result*10 + x%10
		x /= 10
	}
	return (x == result) || (x == result/10)
}
