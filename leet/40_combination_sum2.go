package leet

import (
	"sort"
)

func combinationSum2(candidates []int, target int) (results [][]int) {
	sort.Ints(candidates)
	return cs2(candidates, target)
}

func cs2(candidates []int, target int) (results [][]int) {
	n := len(candidates)
	if n == 0 {
		return nil
	}

	for i := 0; i < n; i++ {
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		if candidates[i] == target {
			results = append(results, []int{candidates[i]})
			continue
		}
		if candidates[i] > target {
			continue
		}

		res := cs2(candidates[i+1:], target-candidates[i])
		if res != nil {
			m := len(res)
			for j := 0; j < m; j++ {
				res[j] = append([]int{candidates[i]}, res[j]...)
			}
			results = append(results, res...)
		}
	}
	return
}
