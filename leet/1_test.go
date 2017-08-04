package leet

import "testing"

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
