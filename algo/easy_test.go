package algo

import (
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
