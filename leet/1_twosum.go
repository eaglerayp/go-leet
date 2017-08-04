package leet

func twoSum(nums []int, target int) []int {
	// key = value, value = index
	m := make(map[int]int)
	for i, v := range nums {
		other := target - v
		if i2, ok := m[other]; ok {
			return []int{i, i2}
		} else {
			m[v] = i
		}
	}
	return nil
}
