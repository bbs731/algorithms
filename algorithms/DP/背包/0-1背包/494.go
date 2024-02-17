package dp

/***
https://www.bilibili.com/video/BV16Y411v7Y6/?vd_source=84c3c489cf545fafdbeb3b3a6cd6a112
这是 0-1 背包的 模版题目。

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


/***
https://leetcode.cn/problems/YaVDxD/solutions/2157241/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-aj9f/

灵神说， 可以优化成一维度的数组

但是我上面那种写法，  dp[.][j] 依赖  j+nums[i],  依赖 j-nums[i] 这个顺序不确定啊， 所以降不了啊。
灵神的解法， j 是单调的，是可以降维的。

 */

