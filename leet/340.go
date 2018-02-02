package leet

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	cm := make(map[byte]int)
	slow := 0
	res := 0
	n := len(s)
	for fast := 0; fast < n; fast++ {
		c := s[fast]
		cm[c] = cm[c] + 1
		for len(cm) > k {
			// move slow to make sure only k chars now
			slowC := s[slow]
			cm[slowC] = cm[slowC] - 1
			if cm[slowC] == 0 {
				delete(cm, slowC)
			}
			slow++
		}
		res = maxInt(res, fast-slow+1)
	}
	return res
}

// method 2
// we always care the last char and last index  to change
