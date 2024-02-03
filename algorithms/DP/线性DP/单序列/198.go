package dp

import "slices"

func rob(nums []int) int {
	n :=len(nums)
	dp := make([]int, n)

	// 为啥，还写的像屎一样。 还有这么多的特判。
	if n <= 2 {
		return slices.Max(nums)
	}

	dp[n-1] = nums[n-1]
	dp[n-2] = max(nums[n-1], nums[n-2])
	for i:=n-3;i>=0; i-- {
		dp[i] = max(dp[i+1], dp[i+2] + nums[i])
	}
	return dp[0]
}


/***
dp[i] = max(dp[i-1], dp[i-2] + nums[i])

相比之下， 这种写法，是不是就简洁的多了。
 */
func rob(nums []int)int {
	n := len(nums)
	dp :=  make([]int, n+2)
	for i:=0; i<n; i++ {
		dp[i+2] = max(dp[i+1], dp[i] + nums[i])
	}
	return dp[n-1+2]
}

func rob(nums []int) int {
	n := len(nums)
	// f0, f1, curr  === >   f1, max(f0+nums[1], f1), curr+1
	f0, f1 := 0 , 0
	for i:=0; i<n; i++ {
		f1, f0 = max(f1, nums[i]+ f0), f1
	}
	return f1
}
