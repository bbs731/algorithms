package dp


/***
从下班的路上就开始想， 想了一整路， 觉得好复杂。
哎， 没想到, 写起来会这么简单。
 */
func maxNonDecreasingLength(nums1 []int, nums2 []int) int {
	n :=len(nums1)
	dp := make([][2]int, n)
	dp[0][0] = 1
	dp[0][1] = 1

	ans := 1
	for i:=1; i<n; i++ {
		dp[i][0]=1
		dp[i][1]=1
		if nums1[i] >= nums1[i-1]{
			dp[i][0] = max(dp[i][0], dp[i-1][0] + 1)
		}
		if nums1[i] >= nums2[i-1]{
			dp[i][0] = max(dp[i][0], dp[i-1][1] + 1)
		}
		if nums2[i] >= nums1[i-1]{
			dp[i][1] = max(dp[i][1], dp[i-1][0] + 1)
		}
		if nums2[i]>= nums2[i-1]{
			dp[i][1] = max(dp[i][1], dp[i-1][1] + 1)
		}
		ans = max(ans, dp[i][0], dp[i][1])
	}
	return ans
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

