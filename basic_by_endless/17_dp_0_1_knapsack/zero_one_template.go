package dp

// 0-1 背包问题的，模板!
func zeroOneKnapsack(values, weights []int, maxW int) int {
	dp := make([]int, maxW+1)

	for i, w := range weights {
		v := values[i]

		for j := maxW; j >= w; j -- {
			dp[j] = max(dp[j], dp[j-w]+v)
		}
	}
	return dp[maxW]
}
