package algo

import "testing"

func TestIsPalindrome(t *testing.T) {
	if isPalindrome("bbab") {
		t.Fail()
	}

	if !isPalindrome("bbbb") {
		t.Fail()
	}

	if isPalindrome("baba") {
		t.Fail()
	}

	if !isPalindrome("bb") {
		t.Fail()
	}

	if !isPalindrome("b") {
		t.Fail()
	}
	if isPalindrome("ba") {
		t.Fail()
	}
}

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
