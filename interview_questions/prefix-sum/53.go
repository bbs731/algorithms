package prefix_sum

/***
https://leetcode.cn/problems/maximum-subarray/solutions/228009/zui-da-zi-xu-he-by-leetcode-solution/
灵神，还给了一个前缀和的解法，真是巧妙啊。
 */

func maxSubArray(nums []int) int {
	min_pre, pre := 0, 0

	ans := nums[0]
	for _, x := range nums {
		pre += x
		// 求 ans, 和 min_pre 的顺序，变的非常的关键。 这个很难一次相对
		ans = max(ans, pre-min_pre)
		min_pre = min(min_pre, pre)
	}
	return ans
}

/***
下面是DP的解法。
 */

func maxSubArray(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	f[0] = nums[0]

	ans := f[0]
	for i := 1; i < n; i++ {
		f[i] = max(f[i-1]+nums[i], nums[i])
		ans = max(ans, f[i])
	}
	return ans
}
