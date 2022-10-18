package backtracking

func generateParenthesis(n int) []string {
	combinations := []string{}
	backtrack(&combinations, "", 0, 0, n)
	return combinations
}

func backtrack(combinations *[]string, str string, open int, close int, n int) {
	if len(str) == 2*n {
		*combinations = append(*combinations, str)
		return
	}

	if open < n {
		backtrack(combinations, str+"(", open+1, close, n)
	}

	if close < open {
		backtrack(combinations, str+")", open, close+1, n)
	}
}

// https://leetcode.com/problems/generate-parentheses/solution/
