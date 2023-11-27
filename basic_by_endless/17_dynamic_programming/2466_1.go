package dp

func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	dp[0] = 1

	for i := 1; i < high+1; i++ {
		//if i >= zero && i >= one {  // this condition is wrong
		if i >= zero {
			dp[i] += dp[i-zero]
		}
		if i >= one {
			dp[i] += dp[i-one]
		}
		dp[i] %= 1e9 + 7
		//}
	}

	ans := 0
	for j := low; j <= high; j++ {
		ans += dp[j]
		ans %= 1e9 + 7
	}
	return ans
}
