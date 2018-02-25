package intowow

func FindListLength(A []int) int {
	// write your code in Go 1.4
	n := len(A)
	if n < 1 {
		return 0
	}
	// iterate the list until we find -1 ?
	i := 0
	count := 0
	for i != -1 {
		i = A[i]
		count++
	}
	return count
}
