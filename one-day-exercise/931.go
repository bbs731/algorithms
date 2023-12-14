package one_day_exercise

/*

https://leetcode.cn/problems/minimum-falling-path-sum/solutions/2341851/cong-di-gui-dao-di-tui-jiao-ni-yi-bu-bu-2cwkb/
灵神的题解

 	f[i][j]= min(f[i+1][j-1], f[i+1][j], f[i+1][j+1])  + matrix[i][j]  // i是倒序的。
	可以降维
 	f[j] = min(f[j-1], f[j], f[j+1]) 需要 tmp variable save previous f[j]
 */
func minFallingPathSum(matrix [][]int) int {
	m := len(matrix[0]) // num of columns
	n := len(matrix)    // num of rows
	inf := int(1e9)

	f := make([]int, m)
	ans := inf

	for i := n - 1; i >= 0; i-- {
		prev := inf // use prev to save f[i+1][j]
		for j := 0; j < m; j++ {
			b := inf
			if j < m-1 {
				b = f[j+1]
			}
			//f[i][j] = min(f[i+1][j], a, b) + matrix[i][j]
			//prev, f[j] = f[j], min(f[j], prev, b)+matrix[i][j] // 我靠， 这个降维，太烧脑了！
			a := f[j]
			f[j] = min(f[j], prev, b) + matrix[i][j]
			prev = a

			if i == 0 {
				ans = min(ans, f[j])
			}
		}
	}
	return ans
}

/*
 	f[i][j]= min(f[i+1][j-1], f[i+1][j], f[i+1][j+1])  + matrix[i][j]  // i是倒序的。
 */
func minFallingPathSum_dp(matrix [][]int) int {
	m := len(matrix[0]) // num of columns
	n := len(matrix)    // num of rows
	inf := int(1e9)

	f := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		f [i] = make([]int, m)
	}
	ans := int(1e9)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < m; j++ {
			a, b := inf, inf
			if j > 0 {
				a = f[i+1][j-1]
			}
			if j < m-1 {
				b = f[i+1][j+1]
			}
			f[i][j] = min(f[i+1][j], a, b) + matrix[i][j]
			if i == 0 {
				ans = min(ans, f[i][j])
			}
		}
	}
	return ans
}

// dfs 的版本会超时！
func minFallingPathSum_dfs(matrix [][]int) int {
	m := len(matrix[0]) // num of columns
	n := len(matrix)    // num of rows
	inf := int(1e9)

	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
		for j := 0; j < m; j++ {
			cache[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == n {
			return 0
		}
		if j == -1 || j == m {
			return inf
		}

		if cache[i][j] != -1 {
			return cache[i][j]
		}
		ans := min(dfs(i+1, j), dfs(i+1, j-1), dfs(i+1, j+1)) + matrix[i][j]
		cache[i][j] = ans
		return ans
	}

	ans := inf
	for j := 0; j < m; j++ {
		ans = min(ans, dfs(0, j))
	}
	return ans

}
