package dp

import "sort"

/*
能做出来这道题， 你需要首先知道  LIS O(n*logn) 的解法。看LC300 的答案。
 */
func LIS(nums []int) []int {
	dp := make([]int, len(nums))
	l := 0
	for i, x := range nums {
		dp[i] = 1
		pos := sort.SearchInts(nums[:l], x) // 这里是 lower bound 因为需要的严格递增的队列。
		nums[pos] = x
		if pos == l {
			l++
		}
		dp[i] = pos + 1
	}
	return dp
}

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	reverse := make([]int, n)

	for i := 0; i < n; i++ {
		reverse[n-1-i] = nums[i]
	}
	dp := LIS(nums)
	rdp := LIS(reverse)

	ans := n + 1
	for i := 0; i < n; i++ {
		if dp[i]+rdp[n-1-i]-1 >= 3 && dp[i] > 1 && rdp[n-1-i] > 1 { // ans 的条件很重要， 错都是错在这儿了。
			ans = min(ans, n-dp[i]-rdp[n-1-i]+1)
		}
	}
	return ans
}
