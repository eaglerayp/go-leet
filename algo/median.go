package algo

import "sort"

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

func isPalindromeIndex(s string, begin, end int) bool {
	if begin < 0 {
		return false
	}
	for begin < end {
		if s[begin] != s[end] {
			return false
		}
		begin++
		end--
	}
	return true
}

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

// #6 Zigzag
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	slices := make([][]rune, numRows)
	for i, _ := range slices {
		slices[i] = []rune{}
	}
	adding := 1
	zindex := 0
	top := numRows - 1
	for _, v := range s {
		if zindex == 0 {
			adding = 1
		} else if zindex == top {
			adding = -1
		}
		slices[zindex] = append(slices[zindex], v)
		zindex += adding
	}
	var result []rune
	for _, v := range slices {
		result = append(result, v...)
	}
	return string(result)
}

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
			} else {
				return min
			}
		}
	}

	return result * sign
}

func maxArea(height []int) int {
	max := 0
	n := len(height)
	if n < 2 {
		return max
	}

	// first O(n) version

	// for i, v := range height {
	// 	width, j := 0, i
	// 	for j < n {
	// 		h := minInt(v, height[j])
	// 		area := h * width
	// 		max = maxInt(max, area)
	// 		width++
	// 		j = i + width
	// 	}
	// }

	// O(logn) version
	i, j := 0, n-1
	for i < j {
		h := minInt(height[i], height[j])
		max = maxInt(max, (j-i)*h)
		for height[i] <= h && i < j {
			i++
		}
		for height[j] <= h && i < j {
			j--
		}
	}

	return max
}

func intToRoman(num int) string {
	// 0<num<4000
	if num < 1 || num > 3999 {
		return ""
	}
	M := []string{"", "M", "MM", "MMM"}
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[(num%10)]
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	result := [][]int{}
	if n < 3 {
		return result
	}
	sort.Ints(nums)
	i := 0
	for i < n-2 {
		if nums[i] > 0 {
			return result
		}
		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				// skip same result
				k--
				j++
				for nums[k] == nums[k+1] && j < k {
					k--
				}
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			} else if sum > 0 {
				// too big
				k--
			} else if sum < 0 {
				// too small
				j++
			}
		}
		// skip same i
		i++
		for i < n-2 && nums[i-1] == nums[i] {
			i++
		}
	}
	return result
}
