package intowow

import (
	"sort"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func FuelQueue(A []int, X int, Y int, Z int) int {
	// write your code in Go 1.4

	// O(n)
	ALiter := [3]int{X, Y, Z}
	ASec := [3]int{0, 0, 0}
	now := 0
	n := len(A)
	i := 0
	for i < n {
		add := false
		for j := 0; j < 3; j++ {
			if ASec[j] <= now && ALiter[j] >= A[i] {
				ASec[j] += A[i]
				ALiter[j] -= A[i]
				add = true
				break
			}
		}
		if add {
			i++
			continue
		}

		// available dispenster not enough
		copy := []int{}
		copy = append(copy, ASec[0])
		copy = append(copy, ASec[1])
		copy = append(copy, ASec[2])
		sort.Ints(copy)
		// if all dispenster availble, return -1
		if now >= copy[2] {
			return -1
		}
		// queue case
		// move now to min(), which > now
		next := now
		for j := 0; next <= now; j++ {
			next = copy[j]
		}
		now = next
		for j := 0; j < 3; j++ {
			if ASec[j] < now {
				ASec[j] = now
			}
		}
	}
	return now
}
