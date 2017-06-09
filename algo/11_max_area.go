package algo

func maxArea(height []int) int {
	max := 0
	n := len(height)
	if n < 2 {
		return max
	}

	// first O(n) version

	// for i, v := range height {
	// 	width, j := 0, i
	// 	for j < n {
	// 		h := minInt(v, height[j])
	// 		area := h * width
	// 		max = maxInt(max, area)
	// 		width++
	// 		j = i + width
	// 	}
	// }

	// O(logn) version
	i, j := 0, n-1
	for i < j {
		h := minInt(height[i], height[j])
		max = maxInt(max, (j-i)*h)
		for height[i] <= h && i < j {
			i++
		}
		for height[j] <= h && i < j {
			j--
		}
	}

	return max
}
