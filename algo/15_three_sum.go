package algo

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	result := [][]int{}
	if n < 3 {
		return result
	}
	sort.Ints(nums)
	i := 0
	for i < n-2 {
		if nums[i] > 0 {
			return result
		}
		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				// skip same result
				k--
				j++
				for nums[k] == nums[k+1] && j < k {
					k--
				}
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			} else if sum > 0 {
				// too big
				k--
			} else if sum < 0 {
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
	return result
}
