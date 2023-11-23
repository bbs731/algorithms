package backtracking

// (从i 开始的逐步减少的 d 个数)  sum of [i-d+1...i] = (i-d+1 + i) * d /2
func combinationSum3_2(k int, n int) [][]int {
	ans := [][]int{}
	path := []int{}

	var dfs func(int, int, int)

	dfs = func(i int, d int, sum int) {
		//剪枝1
		if sum < 0 {
			return
		}
		//剪枝2
		if sum > (2*i-d+1)*d/2 {
			return
		}
		// 剪枝3  //因为我们不再枚举 j 了，就 explicit 的列出  i < d 的剪枝条件
		if i < d {
			return
		}

		if d == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		// 不选
		dfs(i-1, d, sum)

		//选 ith number
		path = append(path, i)
		dfs(i-1, d-1, sum-i)
		path = path[:len(path)-1]
	}

	dfs(9, k, n)
	return ans
}
