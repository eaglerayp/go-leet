package leet

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	if right == 0 {
		return 0
	}
	var mi, result int
	for left < right {
		mi = (left + right) / 2
		if target == nums[mi] {
			return mi
		}
		if target > nums[mi] {
			left = mi + 1
		} else {
			right = mi
		}
	}
	// find insert position
	if target < nums[mi] {
		result = mi
	} else {
		result = mi + 1
	}
	return result
}
