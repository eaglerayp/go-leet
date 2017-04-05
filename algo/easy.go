package algo

type ListNode struct {
	Val  int
	Next *ListNode
}

func twoSum(nums []int, target int) []int {
	// key = value, value = index
	m := make(map[int]int)
	for i, v := range nums {
		other := target - v
		if i2, ok := m[other]; ok {
			return []int{i, i2}
		} else {
			m[v] = i
		}
	}
	return nil
}

func listToNumber(l1 *ListNode) int {
	var i1 int
	d1 := 1
	for l1 != nil {
		i1 = i1 + d1*l1.Val
		l1 = l1.Next
		d1 *= 10
	}
	return i1
}

// You are given **two non-empty** linked lists representing two non-negative integers.
// The digits are stored in reverse order and each of their nodes contain a single digit.
// Add the two numbers and return it as a linked list.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// first version i choose to change it to int then addTwoNumbers
	// but this question should add digit by digit
	null1 := l1 == nil
	null2 := l2 == nil
	carry := 10
	// this is head (dummy node)
	result := &ListNode{}
	prev := result
	ac := 0
	// from digit one to last
	for !null1 && !null2 {
		digit := l1.Val + l2.Val + ac
		if digit >= carry {
			ac = 1
			digit -= carry
		} else {
			ac = 0
		}
		now := &ListNode{Val: digit}
		prev.Next = now
		l1 = l1.Next
		l2 = l2.Next
		prev = now
		null1 = l1 == nil
		null2 = l2 == nil
	}
	if null1 {
		// traverse l2
		for !null2 {
			digit := l2.Val + ac
			if digit >= carry {
				ac = 1
				digit -= carry
			} else {
				ac = 0
			}
			l2.Val = digit
			prev.Next = l2
			prev = l2
			l2 = l2.Next

			null2 = l2 == nil
		}
		// final add carry
		if ac > 0 {
			now := &ListNode{Val: 1}
			prev.Next = now
		}
	} else if null2 {
		// traverse l1
		for !null1 {
			digit := l1.Val + ac
			if digit >= carry {
				ac = 1
				digit -= carry
			} else {
				ac = 0
			}
			l1.Val = digit
			prev.Next = l1
			prev = l1
			l1 = l1.Next
			null1 = l1 == nil
		}
		// final add carry
		if ac > 0 {
			now := &ListNode{Val: 1}
			prev.Next = now
		}
	}
	return result.Next
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

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
	for i, _ := range charMap {
		charMap[i] = -1
	}
	for i, c := range s {
		maxIndex = maxInt(charMap[c]+1, maxIndex)
		charMap[c] = i
		max = maxInt(max, i-maxIndex+1)
	}
	return max
}

func isEven(a int) bool {
	return (a & 0x01) == 0
}

func findMedian(nums []int) float64 {
	n := len(nums)
	if n == 0 {
		return 0.0
	}
	half := n / 2

	if isEven(n) {
		return float64(nums[half-1]+nums[half]) / 2
	}
	return float64(nums[half])
}

// func FindMedianSortedArrays(nums1 []int, nums2 []int) (result float64) {
// 	return findMedianSortedArrays(nums1, nums2)
// }

func reverse(x int) int {
	result := 0
	base := 10
	max := 2147483647
	// use x != 0 can solve negative case
	for x != 0 {
		result = result*base + x%base
		x /= base
	}
	if result > max || result < max*-1 {
		return 0
	}
	return result
}
