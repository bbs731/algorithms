package dp

/***
f[i] = max(f[i+s] + nums[i])  for  i+1<=s <= i+k   // 这个状态第一次还写错了， 真是不应该啊。

朴素的解法，会超时 TLE, 需要优化DP
 */
func maxResult(nums []int, k int) int {
	n := len(nums)
	inf := int(1e9) * n
	f := make([]int, n+1)

	f[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		f[i] = -inf
		for j := i + 1; j <= i+k && j < n; j++ {
			f[i] = max(f[i], f[j]+nums[i])
		}
	}
	return f[0]
}

/**
https://leetcode.cn/problems/jump-game-vi/solutions/2631981/yi-bu-bu-you-hua-cong-di-gui-dao-di-tui-84qn3/
 */
func maxResult(nums []int, k int) int {
	n := len(nums)
	f := make([]int, n+1)

	f[n-1] = nums[n-1]
	// 单调队列
	q := []int{n - 1}
	for i := n - 2; i >= 0; i-- {
		// index out of range, pop out
		for len(q) > 0 && q[0] > i+k {
			q = q[1:]
		}
		f[i] = f[q[0]] + nums[i]
		for len(q) > 0 && f[q[len(q)-1]] <= f[i] {
			// trim tail
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	return f[0]
}
