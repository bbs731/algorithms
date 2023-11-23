package backtracking

// (从i 开始的逐步减少的 d 个数)  sum of [i-d+1...i] = (i-d+1 + i) * d /2
func combinationSum3(k int, n int) [][]int {
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

		//剪枝3  //剪枝3 放到 loop j 的条件里面了
		//if i < d {
		//	return
		//}

		if d == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		// 思路2: 枚举可以选择的值
		for j := i; j > d-1; j-- {
			path = append(path, j)
			dfs(j-1, d-1, sum-j)
			path = path[:len(path)-1]
		}
	}

	dfs(9, k, n)
	return ans
}
