package leet

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	if maxArea([]int{1, 2, 3, 4, 5}) != 6 {
		fmt.Println("T1")
		t.Fail()
	}
	if maxArea([]int{1}) != 0 {
		fmt.Println("T2")
		t.Fail()
	}
	if maxArea([]int{3, 3}) != 3 {
		fmt.Println("T3")
		t.Fail()
	}
}
