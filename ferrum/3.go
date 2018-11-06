package ferrum

// sum of min unique array: inc duplicate element
func getMinimumUniqueSum(arr []int32) int32 {
	// Write your code here
	dup := map[int32]bool{}
	var sum int32 = 0
	for _, v := range arr {
		// duplicate case, inc
		_, ok := dup[v]
		for ok {
			v = v + 1
			_, ok = dup[v]
		}
		dup[v] = true
		sum = sum + v
	}

	return sum
}
