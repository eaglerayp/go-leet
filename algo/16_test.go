package algo

import (
	"fmt"
	"testing"
)

func TestThreeSumCloset(t *testing.T) {
	if threeSumClosest([]int{-1, 2, 1, -4}, 1) != 2 {
		fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
		fmt.Println("t1 fail")
		t.Fail()
	}
	if threeSumClosest([]int{-1, 0, 2}, 1) != 1 {
		fmt.Println(threeSumClosest([]int{-1, 0, 2}, 1))
		fmt.Println("t2 fail")
		t.Fail()
	}
	if threeSumClosest([]int{1, 2, 3, 5, 6, 9}, 6) != 6 {
		fmt.Println(threeSumClosest([]int{1, 2, 3, 5, 6, 9}, 6))
		fmt.Println("t3 fail")
		t.Fail()
	}
}
