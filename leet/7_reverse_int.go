package leet

// func FindMedianSortedArrays(nums1 []int, nums2 []int) (result float64) {
// 	return findMedianSortedArrays(nums1, nums2)
// }

func reverse(x int) int {
	result := 0
	base := 10
	max := 2147483647
	min := -2147483648
	// use x != 0 can solve negative case
	for x != 0 {
		result = result*base + x%base
		x /= base
	}
	if result > max || result < min {
		return 0
	}
	return result
}
