package backtracking

import "strings"

func generateParenthesis(n int) []string {
	ans := []string{}
	path := []string{}

	var dfs func(int, int)

	// open 代表 左括号的个数
	dfs = func(i int, open int) {
		if i == 2*n {
			ans = append(ans, strings.Join(path, ""))
			return
		}

		//选择 '('
		if open < n {
			path = append(path, "(")
			dfs(i+1, open+1)
			path = path[:len(path)-1]
		}

		// 不选择 (
		if open > i-open {
			path = append(path, ")")
			dfs(i+1, open)
			path = path[:len(path)-1]
		}
	}

	dfs(0, 0)
	return ans
}
