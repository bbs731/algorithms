package backtracking

func combine2(n int, k int) (ans [][]int) {
	path := []int{}
	var dfs func(int, int)

	dfs = func(i int, d int) {
		if d == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		// 剪枝
		if i < d {
			return
		}
		// 不选 i
		dfs(i-1, d)

		//选i
		path = append(path, i)
		dfs(i-1, d-1)
		path = path[:len(path)-1]
	}

	dfs(n, k)
	return
}
