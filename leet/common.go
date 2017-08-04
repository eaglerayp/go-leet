package leet

import "fmt"

const (
	MAXINT = 1<<63 - 1
)

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

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0 // return correctly abs(-0)
	}
	return x
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

func iterateList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
