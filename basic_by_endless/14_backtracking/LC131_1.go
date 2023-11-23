package backtracking

func partition(s string) (ans [][]string) {

	path := []string{}
	var dfs func(int)

	n := len(s)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), path...))
			return
		}

		// 这里是在枚举结果
		for j := i; j < n; j++ {
			p := s[i : j+1]
			if isSemetric(p) {
				path = append(path, p)
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return
}

func isSemetric(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
