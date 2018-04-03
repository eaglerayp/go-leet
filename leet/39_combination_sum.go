package leet

func combinationSum(candidates []int, target int) (results [][]int) {
	n := len(candidates)
	if n == 0 {
		return nil
	}

	for i := 0; i < n; i++ {
		if candidates[i] == target {
			results = append(results, []int{candidates[i]})
			continue
		}
		if candidates[i] > target {
			continue
		}

		res := combinationSum(candidates[i:], target-candidates[i])
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
