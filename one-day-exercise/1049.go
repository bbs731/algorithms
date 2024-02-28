package one_day_exercise

func lastStoneWeightII(stones []int) int {
	//n := len(stones)
	sum := 0
	for _, s := range stones {
		sum += s
	}
	target := sum >> 1
	dp := make([]int, target+1)

	// dp[i][j] = max(dp[i][j], dp[i-1][j-stones[i]] + stones[i])
	for _, s := range stones {
		for j := target; j >= 0; j-- {
			dp[j] = max(dp[j], dp[j-s]+s)
		}
	}
	return sum - 2*dp[target]
}
