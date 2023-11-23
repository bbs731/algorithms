package backtracking

func partition2(s string) (ans [][]string) {
	path := []string{}
	var dfs func(int, int)

	n := len(s)
	dfs = func(i int, pos int) {
		if i == n {
			if pos == n {
				ans = append(ans, append([]string(nil), path...))
			}
			return
		}
		// 不选
		dfs(i+1, pos)

		// 选
		p := s[pos : i+1]
		if isSemetric(p) {
			path = append(path, p)
			dfs(i+1, i+1)
			path = path[:len(path)-1]
		}
	}

	dfs(0, 0)
	return
}
