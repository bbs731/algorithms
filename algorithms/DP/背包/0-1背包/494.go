package dp

/***
春雷， 你这个思路也是够野的啊！

	dp[i][sum] = dp[i+1][sum+ nums[i]] + dp[i+1][sum - nums[i]]
 */
func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	dp := make([][2001]int, n)

	// 如果 nums[n-1] = 0 有 bug
	//dp[n-1][1000+nums[n-1]] = 1
	//dp[n-1][1000-nums[n-1]] = 1
	dp[n-1][1000+nums[n-1]] += 1
	dp[n-1][1000-nums[n-1]] += 1


	for i:=n-2; i>=0; i-- {
		for j:=nums[i]; j<2001-nums[i]; j++ {
			dp[i][j] = dp[i+1][j+nums[i]] + dp[i+1][j-nums[i]]
		}
	}
	return dp[0][1000+target]
}
