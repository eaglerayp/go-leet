package algo

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	ex1 := []int{1, 3}
	ex2 := []int{2}
	ex3 := []int{1, 2}
	ex4 := []int{3, 4}
	ex5 := []int{3}
	ex6 := []int{1, 2, 3}
	ex7 := []int{1}
	ex8 := []int{2, 3, 4}
	a1 := findMedianSortedArrays(ex1, ex2)
	a2 := findMedianSortedArrays(ex3, ex4)
	a3 := findMedianSortedArrays(ex1, ex3)
	a4 := findMedianSortedArrays(ex1, ex4)
	a5 := findMedianSortedArrays(ex3, ex5)
	a6 := findMedianSortedArrays(ex1, []int{})
	a7 := findMedianSortedArrays(ex3, ex6)
	a8 := findMedianSortedArrays(ex7, ex8)
	if a1 != 2 {
		fmt.Println("a1")
		t.Fail()
	}

	if a2 != 2.5 {
		fmt.Println("a2")
		t.Fail()
	}

	if a3 != 1.5 {
		fmt.Println("a3")
		t.Fail()
	}

	if a4 != 3 {
		fmt.Println("a4")
		t.Fail()
	}

	if a5 != 2 {
		fmt.Println("a5")
		t.Fail()
	}

	if a6 != 2 {
		fmt.Println("a6")
		t.Fail()
	}

	if a7 != 2 {
		fmt.Println("a7")
		t.Fail()
	}

	if a8 != 2.5 {
		fmt.Println("a8")
		t.Fail()
	}

	if findMedianSortedArrays(ex2, []int{1, 3, 4}) != 2.5 {
		fmt.Println("a9")
		t.Fail()
	}

	if findMedianSortedArrays(ex2, []int{1, 3, 4, 5, 6}) != 3.5 {
		fmt.Println("a10")
		t.Fail()
	}

	if findMedianSortedArrays([]int{7, 8}, []int{1, 2, 3, 4, 5, 6, 9}) != 5 {
		fmt.Println("a11")
		t.Fail()
	}

	if findMedianSortedArrays([]int{1, 3, 4, 5}, []int{2}) != 3 {
		fmt.Println("a12")
		t.Fail()
	}

	if findMedianSortedArrays([]int{3, 4}, []int{1, 2, 5}) != 3 {
		fmt.Println("a13")
		t.Fail()
	}

	if findMedianSortedArrays([]int{1, 4}, []int{2, 3, 5}) != 3 {
		fmt.Println("a14")
		t.Fail()
	}
}
