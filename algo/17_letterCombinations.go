package algo

func letterCombinations(digits string) []string {
	n := len(digits)
	result := []string{}
	if n == 0 {
		return result
	}
	dict := make(map[string]string)
	dict["0"] = ""
	dict["1"] = ""
	dict["2"] = "abc"
	dict["3"] = "def"
	dict["4"] = "ghi"
	dict["5"] = "jkl"
	dict["6"] = "mno"
	dict["7"] = "pqrs"
	dict["8"] = "tuv"
	dict["9"] = "wxyz"
	result = append(result, "")
	for i := 0; i < n; i++ {
		temp := []string{}
		v := dict[string(digits[i])]
		tempN := len(result)
		for _, newC := range v {
			// pop result and  push temp
			for j := 0; j < tempN; j++ {
				temp = append(temp, result[j]+string(newC))
			}
		}
		result = temp
	}
	return result
}
