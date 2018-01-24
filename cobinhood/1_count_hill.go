package cobinhood

const (
	INIT = 0
	DOWN = -1
	UP   = 1
)

func Solution(A []int) int {
	// write your code in Go 1.4

	n := len(A)
	if n < 1 {
		return 0
	}
	solution := 1

	// i think better soluition with O(n) space memeory
	// diff := []bool{}
	// for i := 1; i < n; i++ {
	// 	last := A[i-1]
	// 	if last != A[i] {
	// 		diff = append(diff, A[i] > last)
	// 	}
	// }
	// if len(diff) == 0 {
	// 	return solution
	// }
	// // the last one
	// solution++
	// n = len(diff)
	// for i := 1; i < n; i++ {
	// 	if diff[i-1] != diff[i] {
	// 		solution++
	// 	}
	// }

	// init sign,  0 == int, -1 downside, 1 upside
	sign := 0
	for i := 1; i < n; i++ {
		last := A[i-1]
		if last > A[i] {
			// downside
			if sign != DOWN {
				sign = DOWN
				solution++
			}
		} else if last < A[i] {
			// upside
			if sign != UP {
				sign = UP
				solution++
			}
		}
	}
	return solution
}
