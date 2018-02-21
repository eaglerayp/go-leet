package leet

func searchRange(nums []int, target int) []int {
	n := len(nums)
	left := 0
	right := n
	var mi int
	exists := false
	for left < right {
		mi = (left + right) / 2
		if target == nums[mi] {
			exists = true
			break
		}
		if target > nums[mi] {
			left = mi + 1
		} else {
			right = mi
		}
	}
	if !exists {
		return []int{-1, -1}
	}

	// find range by search >= and <= target
	var li, ri int
	// find left index which smaller than target
	left = 0
	right = mi
	for left < right {
		li = (left + right) / 2
		if nums[li] < target {
			left = li + 1
		} else {
			right = li
		}
	}
	if nums[left] == target {
		li = left
	}

	// find right index which bigger than target
	left = mi
	right = n
	for left < right {
		ri = (left + right) / 2
		if nums[ri] > target {
			right = ri
		} else {
			left = ri + 1
		}
	}
	if nums[right-1] == target {
		ri = right - 1
	}
	return []int{li, ri}
}
