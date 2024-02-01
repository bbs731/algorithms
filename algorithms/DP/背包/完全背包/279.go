package 完全背包

/***
 dp[n] = min(dp[n] ,  dp[n-i*i] + 1)  // for all i that i*i <=n


12 = 4 + 4 + 4
因为 1个数，可以选任意次数，所以是完全背包问题。
 */
func numSquares(n int) int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}
