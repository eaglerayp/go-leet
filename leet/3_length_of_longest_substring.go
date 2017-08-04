package leet

func lengthOfLongestSubstring(s string) int {
	// Old version
	// // traverse the string
	// // keep the longest substrings' character
	// maxLen := 0
	// nowLen := 0
	// maxC := make(map[rune]bool)
	// startIndex := 0
	// // maxIndex := 0

	// for i, char := range s {
	// 	if maxC[char] {
	// 		// remove the character, and len begin from that
	// 		for startIndex < i {
	// 			nc := rune(s[startIndex])
	// 			startIndex++
	// 			if nc == char {
	// 				break
	// 			}
	// 			nowLen--
	// 			delete(maxC, nc)
	// 		}
	// 		continue
	// 	}
	// 	// extend the substring
	// 	nowLen++
	// 	maxC[char] = true
	// 	if nowLen > maxLen {
	// 		maxLen = nowLen
	// 		// maxIndex = startIndex
	// 	}
	// }
	// return maxLen

	// refine solution with DP
	max, maxIndex := 0, 0
	charMap := make([]int, 256)
	for i := range charMap {
		charMap[i] = -1
	}
	for i, c := range s {
		maxIndex = maxInt(charMap[c]+1, maxIndex)
		charMap[c] = i
		max = maxInt(max, i-maxIndex+1)
	}
	return max
}
