package leet

import "sort"

func threeSumClosest(nums []int, target int) (ans int) {
	ans = 0
	minDiff := MAXINT
	n := len(nums)
	if n < 3 {
		return 0
	}

	sort.Ints(nums)
	i := 0
	for i < n-2 {
		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			diff := sum - target
			absDiff := absInt(diff)
			if absDiff < minDiff {
				minDiff = absDiff
				ans = sum
			}
			if diff == 0 {
				return ans
			} else if diff > 0 {
				// too big
				k--
			} else if diff < 0 {
				// too small
				j++
			}
		}
		// skip same i
		i++
		for i < n-2 && nums[i-1] == nums[i] {
			i++
		}
	}
	return
}
