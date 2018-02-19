package leet

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	// first we find pivot (smallest number)
	left := 0
	right := n - 1
	pivot := 0
	for left < right {
		pivot = (left + right) / 2
		if target == nums[pivot] {
			return pivot
		}
		if nums[pivot] > nums[right] {
			left = pivot + 1
		} else {
			right = pivot
		}
	}
	if target == nums[n-1] {
		return n - 1
	}
	if target > nums[n-1] {
		right = pivot
		left = 0
	} else {
		left = pivot
		right = n - 1
	}
	mi := 0
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
	return -1
}
