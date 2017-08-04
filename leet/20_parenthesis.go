package leet

var parenthesisMap = map[rune]rune{
	'}': '{',
	')': '(',
	']': '[',
}

func isValid(s string) bool {
	arr := []rune{}
	i := 0
	// {()} is correct?
	for _, c := range s {
		switch c {
		case '{', '(', '[':
			if i >= len(arr) {
				arr = append(arr, c)
				i++
				continue
			}
			arr[i] = c
			i++
		// add to stack
		case '}', ')', ']':
			if i < 1 {
				return false
			}
			if arr[i-1] != parenthesisMap[c] {
				return false
			}
			i--
		}
	}
	return i == 0
}
