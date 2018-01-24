package cobinhood

import (
	"sort"
)

const (
	MAX_INT = 1<<63 - 1
	MIN_INT = -1 << 63
)

func AdjacentMinimumDistance(A []int) int {
	// find adjacent pair between A means that to find no-cotinuous elements's distance

	sort.Ints(A)
	n := len(A)
	if n == 0 {
		return -1
	}
	// corner case  diff > MAX_INT
	dist := MAX_INT
	for i := 1; i < n; i++ {
		last := A[i-1]
		cur := A[i]
		// may overflow here
		diff := cur - last
		if diff == 0 {
			return 0
		} else if diff == 1 {
			// not adjacent case
		} else {
			if diff < dist {
				dist = diff
			}
		}
	}

	// no adjacent case
	if dist == MAX_INT {
		return -1
	}

	// // special case 2, overflow case
	if dist > 100000000 || dist < 0 {
		return -2
	}
	return dist
}
