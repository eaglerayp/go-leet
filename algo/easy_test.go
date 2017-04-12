package algo

import (
	"fmt"
	"testing"
)

func sameValueSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	ma := make(map[int]bool)
	for _, v := range a {
		ma[v] = true
	}
	for _, v := range b {
		if _, ok := ma[v]; !ok {
			return false
		}
		delete(ma, v)
	}
	return true
}

func TestTwoSum(t *testing.T) {
	a1 := (twoSum([]int{3, 2, 4}, 6))
	if !sameValueSlice(a1, []int{1, 2}) {
		t.Fail()
	}
	a2 := (twoSum([]int{3, 3}, 6))
	if !sameValueSlice(a2, []int{0, 1}) {
		t.Fail()
	}
}

func TestAddTwoNumbers(t *testing.T) {
	e342 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	e465 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	e999 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}
	e0 := &ListNode{Val: 0}
	e1 := &ListNode{Val: 1}
	a1 := listToNumber(addTwoNumbers(e342, e465))
	a2 := listToNumber(addTwoNumbers(e0, e0))
	a3 := listToNumber(addTwoNumbers(e342, e0))
	a4 := listToNumber(addTwoNumbers(e999, e1))
	// fmt.Println(a1)
	// fmt.Println(a2)
	// fmt.Println(a3)
	// fmt.Println(a4)
	if a1 != 807 {
		t.Fail()
	}
	if a2 != 0 {
		t.Fail()
	}
	if a3 != 342 {
		t.Fail()
	}
	if a4 != 1000 {
		t.Fail()
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	a1 := lengthOfLongestSubstring("bbbbb")
	a2 := lengthOfLongestSubstring("pwwkew")
	a3 := lengthOfLongestSubstring("abcabcbb")
	a4 := lengthOfLongestSubstring("pwwkewa")
	if a1 != 1 {
		t.Fail()
	}

	if a2 != 3 {
		t.Fail()
	}

	if a3 != 3 {
		t.Fail()
	}

	if a4 != 4 {
		t.Fail()
	}
}

func TestReverse(t *testing.T) {
	if reverse(123) != 321 {
		t.Fail()
	}
	// bigger than  2,147,483,647
	if reverse(1111111119) != 0 {
		t.Fail()
	}

	if reverse(0) != 0 {
		t.Fail()
	}

	if reverse(-123456789) != -987654321 {
		t.Fail()
	}
}

func TestIsPalindromeInteger(t *testing.T) {
	if isPalindromeInteger(123) {
		t.Fail()
	}
	// bigger than  2,147,483,647
	if !isPalindromeInteger(1) {
		t.Fail()
	}

	if isPalindromeInteger(-1) {
		t.Fail()
	}

	if !isPalindromeInteger(121) {
		t.Fail()
	}
}

func TestRomanToInt(t *testing.T) {
	if romanToInt("MMMCMXCIX") != 3999 {
		fmt.Println("T1")
		t.Fail()
	}

	if romanToInt("I") != 1 {
		fmt.Println("T2")
		t.Fail()
	}

	if romanToInt("CDXCIX") != 499 {
		fmt.Println("T3")
		t.Fail()
	}
	if romanToInt("") != 0 {
		fmt.Println("T4")
		t.Fail()
	}
}
