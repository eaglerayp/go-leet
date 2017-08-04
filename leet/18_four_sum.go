package leet

import "sort"

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	result := [][]int{}
	if n < 4 {
		return result
	}
	sort.Ints(nums)
	i := 0
	for i < n-3 {
		tempRes := threeSumForFourSum(nums[i+1:], target-nums[i])
		tn := len(tempRes)
		for j := 0; j < tn; j++ {
			result = append(result, append(tempRes[j], nums[i]))
		}
		// skip same i
		i++
		for i < n-3 && nums[i-1] == nums[i] {
			i++
		}
	}
	return result
}

func threeSumForFourSum(nums []int, target int) [][]int {
	n := len(nums)
	result := [][]int{}
	if n < 3 {
		return result
	}
	// input is an sorted array
	i := 0
	for i < n-2 {
		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
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
			} else if sum > target {
				// too big
				k--
			} else if sum < target {
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
