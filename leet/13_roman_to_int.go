package leet

func romanToInt(s string) int {
	n := len(s)
	if n < 1 {
		return 0
	}
	nums := make(map[byte]int)
	nums['I'] = 1
	nums['V'] = 5
	nums['X'] = 10
	nums['L'] = 50
	nums['C'] = 100
	nums['D'] = 500
	nums['M'] = 1000

	result := 0

	for i := 0; i < n-1; i++ {
		if nums[s[i]] < nums[s[i+1]] {
			result -= nums[s[i]]
		} else {
			result += nums[s[i]]
		}
	}
	return result + nums[s[n-1]]
}
