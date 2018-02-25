package intowow

func CrossRiver(A []int, D int) int {
	// write your code in Go 1.4
	n := len(A)
	if D > n {
		return 0
	}
	farIndex := D - 1
	time := 0
	for farIndex < n {
		if time == 100001 {
			return -1
		}
		min := 100001
		jump := false
		canWalk := farIndex
		if farIndex > n-1 {
			canWalk = n - 1
		}

		for i := 0; i <= canWalk; i++ {
			if (A[i] != -1) && (A[i] <= time) {
				if (i + D) > farIndex {
					farIndex = i + D
					jump = true
					continue
				}
			}
			// need to wait time
			if (A[i] > time) && (A[i] < min) {
				min = A[i]
			}
		}
		// log.Println("Time:", time, "Far:", farIndex)
		if jump {
			continue
		}

		// move to next time
		time = min
	}

	return time
}
