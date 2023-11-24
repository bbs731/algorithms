package backtracking

import "strings"

func letterCasePermutation(s string) []string {
	n := len(s)
	var dfs func(int)
	ans := []string{}
	path := []string{}

	dfs = func(i int) {
		if i == n {
			ans = append(ans, strings.Join(path, ""))
			return
		}

		if s[i] >= '0' && s[i] <= '9' {
			path = append(path, string(s[i]))
			dfs(i + 1)
		} else {
			// we have two options here
			path = append(path, string(s[i]))
			dfs(i + 1)
			path = path[:len(path)-1]

			if s[i] >= 'a' {
				path = append(path, string(s[i]-('a'-'A')))
				dfs(i + 1)
			} else {
				path = append(path, string(s[i]+('a'-'A')))
				dfs(i + 1)
			}
		}
		// 这个恢复现场很重要啊！不要忘记
		path = path[:len(path)-1]

	}

	dfs(0)
	return ans
}
