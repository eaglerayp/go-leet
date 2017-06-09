package algo

// Zigzag
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	slices := make([][]rune, numRows)
	for i, _ := range slices {
		slices[i] = []rune{}
	}
	adding := 1
	zindex := 0
	top := numRows - 1
	for _, v := range s {
		if zindex == 0 {
			adding = 1
		} else if zindex == top {
			adding = -1
		}
		slices[zindex] = append(slices[zindex], v)
		zindex += adding
	}
	var result []rune
	for _, v := range slices {
		result = append(result, v...)
	}
	return string(result)
}
