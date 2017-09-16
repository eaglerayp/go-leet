package leet

func generateParenthesis(n int) []string {
	result := []string{}
	addSubParenthesis(&result, "", n, 0)
	return result
}

func addSubParenthesis(list *[]string, str string, left, right int) {
	if left == 0 && right == 0 {
		*list = append(*list, str)
		return
	}
	if right > 0 {
		addSubParenthesis(list, str+")", left, right-1)
	}
	if left > 0 {
		addSubParenthesis(list, str+"(", left-1, right+1)
	}
}
