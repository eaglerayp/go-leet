package leet

func myAtoi(str string) int {
	result := 0
	if len(str) < 1 {
		return 0
	}
	// reach first num or + -
	firstIndex := 0
	n := len(str)
	for firstIndex < n {
		if str[firstIndex] != ' ' {
			break
		}
		firstIndex++
	}
	// check sign
	sign := 1
	if str[firstIndex] == '+' {
		firstIndex++
	} else if str[firstIndex] == '-' {
		sign = -1
		firstIndex++
	}

	chMap := make(map[rune]int)
	for i, v := range "0123456789" {
		chMap[v] = i
	}

	max := 2147483647
	min := -2147483648
	for i := firstIndex; i < n; i++ {
		v := rune(str[i])
		digit, ok := chMap[v]
		if !ok {
			break
		}
		result = result*10 + digit
		if result > max/10 || (result == max/10 && digit-0 > 7) {
			if sign == 1 {
				return max
			}
			return min
		}
	}

	return result * sign
}
