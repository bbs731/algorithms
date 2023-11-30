package dp

/*
	p = nums1[i] * nums2[j]
	f[i+1][j+1] = max( f[i][j], f[i][j] + p , p, f[i][j+1], f[i+1][j])

	f[j+1] = max(f[j], prevj +p, f[j+1], f[j])  need one more variable to save f[i][j] as prevj 因为降维一维数组之后无法区分 f[i][j] and f[i+1][j])
*/

func maxDotProduct(nums1 []int, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)
	inf := int(1e9)

	dp := make([]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = -inf // 这行翻译的是  dfs 中 j < 0  return -inf    DP 中  dp[i][0] = -inf
	}

	for i := 0; i < n; i++ {
		prevj := dp[0] // prevj 初始化为 dp[i][j]
		dp[0] = -inf   // 这行翻译的是  dfs 中 i < 0  return -inf   DP 中  dp[0][j] = -inf
		for j := 0; j < m; j++ {
			tmp := dp[j+1]
			p := nums1[i] * nums2[j]
			dp[j+1] = max(prevj, prevj+p, dp[j+1], dp[j], p)
			prevj = tmp
		}
	}
	return dp[m]
}

/*
	p = nums1[i] * nums2[j]
	f[i+1][j+1] = max( f[i][j], f[i][j] + p , p, f[i][j+1], f[i+1][j])
 */
func maxDotProduct_dp(nums1 []int, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)
	inf := int(1e9)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		dp[i][0] = -inf
		for j := 0; j <= m; j++ {
			dp[0][j] = -inf
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			p := nums1[i] * nums2[j]
			dp[i+1][j+1] = max(dp[i][j], dp[i][j]+p, dp[i][j+1], dp[i+1][j], p)
		}
	}
	return dp[n][m]
}

func maxDotProduct_dfs(nums1 []int, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)
	var dfs func(int, int) int
	inf := int(1e9)

	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
		for j := 0; j < m; j++ {
			cache[i][j] = inf
		}
	}

	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return -inf
		}

		if cache[i][j] != inf {
			return cache[i][j]
		}

		p := nums1[i] * nums2[j]
		ans := max(dfs(i-1, j-1), dfs(i-1, j-1)+p, dfs(i-1, j), dfs(i, j-1), p)

		cache[i][j] = ans
		return ans
	}
	return dfs(n-1, m-1)
}
