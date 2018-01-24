package leet

func findSubstring(s string, words []string) []int {
	// first idea: compose all substring cases, but this one fail, is costs too much memory

	// just records words and count, then try all substring
	result := []int{}
	if len(words) == 0 || len(s) == 0 {
		return result
	}
	wordsMap := make(map[string]int)
	n := len(words)
	wordlen := len(words[0])
	strLen := wordlen * n
	for i := 0; i < n; i++ {
		wordsMap[words[i]]++
	}

	maxStart := len(s) - strLen
	if maxStart < 0 {
		return result
	}

	for i := 0; i <= maxStart; i++ {
		// copy wordsmap
		copyMap := make(map[string]int)
		for k, v := range wordsMap {
			copyMap[k] = v
		}
		// try all words
		for j := 0; j < n; j++ {
			sub := s[i+j*wordlen : i+(j+1)*wordlen]
			v := copyMap[sub]
			if v == 0 {
				break
			} else if v == 1 {
				delete(copyMap, sub)
			} else {
				copyMap[sub] = v - 1
			}
		}
		if len(copyMap) == 0 {
			result = append(result, i)
		}
	}
	return result
}
