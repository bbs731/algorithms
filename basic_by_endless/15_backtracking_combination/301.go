package backtracking

import "strings"

func removeInvalidParentheses(s string) []string {
	ans := []string{}
	path := []string{}
	n := len(s)
	var dfs func(int, int, int)

	dfs = func(i int, left, right int) {
		if i == n {
			// validation
			if left == right {
				if len(ans) == 0 || len(ans[0]) == len(path) {
					ans = append(ans, strings.Join(path, ""))
				} else if len(path) > len(ans[0]) {
					ans = append([]string(nil), strings.Join(path, ""))
				} else {
					// do nothing
				}
			}
			return
		}

		if left < right {
			return
		}

		if s[i] == '(' {
			// select
			path = append(path, "(")
			dfs(i+1, left+1, right)
			path = path[:len(path)-1]

			// no select
			dfs(i+1, left, right)

		} else if s[i] == ')' {
			// select
			if left > right {
				path = append(path, ")")
				dfs(i+1, left, right+1)
				path = path[:len(path)-1]
			}

			// no select
			dfs(i+1, left, right)

		} else {
			// alphabet
			path = append(path, string(s[i]))
			dfs(i+1, left, right)
			// 这里需要恢复现场
			path = path[:len(path)-1]
		}
	}

	dfs(0, 0, 0)

	//ToDo: 既然有去重的需求，那么，实际上应该有剪枝的可能，不然不会出现重复的数据
	// remove duplicate
	visited := make(map[string]bool)
	l := make([]string, 0)
	for _, a := range ans {
		if _, ok := visited[a]; ok {
			continue
		}
		l = append(l, a)
		visited[a] = true
	}
	return l
}
