package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestBinarySearch(t *testing.T) {
	assert.Equal(t, 0, binarySearch([]int{1, 6, 8, 9, 13}, 1))
	assert.Equal(t, 4, binarySearch([]int{1, 6, 8, 9, 13}, 13))
	assert.Equal(t, 5, binarySearch([]int{1, 6, 8, 9, 10, 13}, 13))
	assert.Equal(t, -1, binarySearch([]int{1, 6, 8, 9, 10, 13}, 7))
	assert.Equal(t, -1, binarySearch([]int{}, 2))
	assert.Equal(t, 0, binarySearch([]int{2}, 2))
	assert.Equal(t, -1, binarySearch([]int{2}, 1))
	assert.Equal(t, 0, binarySearch([]int{0, 1}, 0))
	assert.Equal(t, 1, binarySearch([]int{0, 1}, 1))
	assert.Equal(t, -1, binarySearch([]int{0, 1}, -1))
	assert.Equal(t, 1, binarySearch([]int{1, 4}, 4))
}
