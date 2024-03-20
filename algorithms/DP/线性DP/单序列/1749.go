package dp

/***
真是赏心悦目的逻辑实验！
 */
func maxAbsoluteSum(nums []int) int {
	ans := 0
	n := len(nums)
	//f := make([]int, n+1)
	//g := make([]int, n+1)
	f0 := 0
	g0 := 0

	for i := 1; i <= n; i++ {
		f0 = max(f0+nums[i-1], nums[i-1], 0)
		g0 = min(g0+nums[i-1], nums[i-1], 0)
		//f[i] = max(f[i-1]+nums[i-1], nums[i-1], 0)
		//g[i] = min(g[i-1]+nums[i-1], nums[i-1], 0)
		ans = max(ans, f0, -g0)
	}
	return ans
}
