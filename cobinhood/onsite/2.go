package main

func Circle(relations [][]bool) int {
	result := map[int]int{}
	n := len(relations)
	for i := 0; i < n; i++ {
		result[i] = i
	}
	for i := 0; i < n; i++ {
		relation := relations[i]
		for j := 0; j < n; j++ {
			if relation[j] {
				result[j] = min(result[j], result[i])
			}
		}
	}
	prev := 0
	count := 1
	for _, v := range result {
		if v != prev {
			count = count + 1
			prev = v
		}
	}
	return count
}
