package dp

/*
dfs(i, j) 合并 从 [i..j] stones 需要最小的 cost
 */
func mergeStones(stones []int, k int) int {
	n := len(stones)
	psum := make([]int, n+1)
	inf := int(1e9)

	// sum[i..j] = psum[j+1] - psum[i]
	for i := 1; i <= n; i++ {
		psum[i] = psum[i-1] + stones[i-1]
	}

	if (n-k)%(k-1) != 0 {
		return -1
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i > j {
			return 0
		}
		if j-i+1 <= k {
			return psum[j+1] - psum[i]
		}

		ans := inf
		ans = min(ans, dfs(i, i+k-1)+dfs(i+k, j))
		for q := 0; q < k-1; q++ {
			ans = min(ans, dfs(i, i+q)+dfs(i+q+1, j))
		}

		return ans
	}
	return dfs(0, n-1)
}
