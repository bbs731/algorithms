package dp


/***
视频讲解：
https://www.bilibili.com/video/BV16Y411v7Y6/?vd_source=84c3c489cf545fafdbeb3b3a6cd6a112

dfs(i,c) = max( dfs(i-1,c),  dfs(i-1，c-w[i]) + v[i])
 */
func zero_one_knapsack(capacity int, w []int, v []int) int{
	n := len(w)

	var dfs func (int, int) int
	dfs = func(i, c int) int {
		if i <0 {
			return 0
		}
		if c < w[i] {
			// 不选
			return dfs(i-1, c)
		}

		return max(dfs(i-1, c), dfs(i-1, c-w[i]) + v[i])
	}
	return dfs(n-1, capacity)
}



func max(a, b int) int {
	if  a > b {
		return a
	}
	return b
}
