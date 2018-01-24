package cobinhood

import "fmt"

const (
	DECIMAL_SIX = 1000000
)

func MatrixSquare(base, m1 []int, n int) []int {
	fmt.Println(n, ":", m1)
	result := []int{0, 0, 0, 0}
	if n == 1 {
		return m1
	}
	if (n % 2) == 1 {
		// A * A^n
		m12 := MatrixSquare(base, m1, n-1)
		result[0] = m1[0]*m12[0] + m1[1]*m12[2]
		result[1] = m1[0]*m12[1] + m1[1]*m12[3]
		result[2] = m1[2]*m12[0] + m1[3]*m12[2]
		result[3] = m1[2]*m12[1] + m1[3]*m12[3]
		return result
	}
	// ^2
	result[0] = m1[0]*m1[0] + m1[1]*m1[2]
	result[1] = m1[0]*m1[1] + m1[1]*m1[3]
	result[2] = m1[2]*m1[0] + m1[3]*m1[2]
	result[3] = m1[2]*m1[1] + m1[3]*m1[3]
	return MatrixSquare(base, result, n/2)
}

func Fibonacci(N int) int {
	// compute complexity O(logN)
	if N < 2 {
		return N
	}

	// if we want O(logN),  we cannot use multiply for preserved 6-digigt??
	if N == 2 {
		return 1
	}
	matrix := []int{1, 1, 1, 0}
	result := MatrixSquare(matrix, matrix, N-2)
	return (result[0] + result[1]) % DECIMAL_SIX

}
