package dp

/*

这道题是个地域级别， k=2 时，是一道经典的， 区间DP 问题。
把 k 参数化之后就是个地狱的难度。不过可以好好锻炼一下思维。


dfs(i, j, 1) = dfs(i, j, k) + sum[i..j]  这个想的就比较难。
dfs(i,j,p) = min(dfs(i, m, 1) ... dfs(m+1, j, p-1))  其中  i< m = i+ (k-1)x  <j && p >= 2


dfs(i, j) 合并 从 [i..j] stones 需要最小的 cost, 那么 dfs(i, j) 是怎么来的呢？
由， dfs(i, k) 和 dfs(k+1，j) 转换而来。

dfs(i, j) = min( dfs(i, k) +dfs(k+1, j) ) + sum[i..j]
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
		if i >= j {
			return 0
		}

		if j-i+1 <= k {
			return psum[j+1] - psum[i]
		}

		ans := inf
		for q := i+1; q < j; q++ {
			ans = min(ans, dfs(i, q)+dfs(q+1, j))
		}
		ans += psum[j+1]- psum[i]

		return ans
	}
	return dfs(0, n-1)
}
