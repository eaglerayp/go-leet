package algo

func findMedianSortedArrays(nums1 []int, nums2 []int) (result float64) {
	// design, from middle1 increase to find the balance value
	// make sure len2 > len1

	n1 := len(nums1)
	n2 := len(nums2)

	if n2 < n1 {
		n1, n2, nums1, nums2 = n2, n1, nums2, nums1
	}

	if n2 == 0 {
		return 0.0
	}
	// target i1,i2   (i1+i2)*2 =n1+n2 (+1)
	// i2 = m+n+1/2 - i1
	i1min, i1max := 0, n1

	for i1min <= i1max {
		i1 := (i1min + i1max) / 2
		i2 := (n1+n2+1)/2 - i1
		if i1 < n1 && nums1[i1] < nums2[i2-1] {
			// i too small
			i1min = i1 + 1
		} else if i1 > 0 && nums1[i1-1] > nums2[i2] {
			// i too big
			i1max = i1 - 1
		} else {
			// get i

			// then find max on left part, and min on right part
			var maxLeft, minRight int
			if i1 == 0 {
				maxLeft = nums2[i2-1]
			} else if i2 == 0 {
				maxLeft = nums1[i1-1]
			} else {
				maxLeft = maxInt(nums2[i2-1], nums1[i1-1])
			}

			if !isEven(n1 + n2) {
				return float64(maxLeft)
			}
			if i1 == n1 {
				minRight = nums2[i2]
			} else if i2 == n2 {
				minRight = nums1[i1]
			} else {
				minRight = minInt(nums2[i2], nums1[i1])
			}
			return float64(maxLeft+minRight) / 2
		}
	}

	return
}
