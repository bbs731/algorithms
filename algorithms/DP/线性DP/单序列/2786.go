package dp

/***
[8,50,65,85,8,73,55,50,29,95,5,68,52,79]
 */

func maxScore(nums []int, x int) int64 {
	n := len(nums)
	dp := [2]int{nums[0], nums[0]}
	if nums[0]&1 == 0 {
		dp[1] -= x
	} else {
		dp[0] -= x
	}

	for i := 1; i < n; i++ {
		num := nums[i]
		if num&1 == 0 {
			//dp[0] = nums[i] + max(dp[0], dp[1]-x)
			dp[0] = max(dp[0]+num, dp[1]-x+num)
		} else {
			//dp[1] = nums[i] + max(dp[1], dp[0]-x)
			dp[1] = max(dp[1]+num, dp[0]-x+num)
		}
	}
	return int64(max(dp[0], dp[1]))
}

/***

这道题，十分的困难。 难在了初始化上不是特别好像对！
 */

func maxScore(nums []int, x int) int64 {
	n := len(nums)
	fodd := nums[0]
	feven := nums[0]
	if nums[0]&1 == 0 {
		fodd = -x
	} else {
		feven = -x
	}

	for i := 1; i < n; i++ {
		if nums[i]&1 != 0 {
			// odd
			fodd = max(fodd+nums[i], nums[i]-x+feven)
			feven = max(feven, nums[i]-2*x+feven)

		} else {
			// even
			feven = max(feven+nums[i], nums[i]-x+fodd)
			fodd = max(fodd, nums[i]-2*x+fodd)

		}
	}
	return int64(max(feven, fodd))
}
