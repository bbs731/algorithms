package dp

func maxArrayValue(nums []int) int64 {
	n := len(nums)
	// f[i] 只依赖 f[i+1] 其实，可以用一个变量表示，也是可以的。
	f := make([]int, n+1)
	f[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= f[i+1] {
			f[i] = f[i+1] + nums[i]
		} else {
			f[i] = nums[i]
		}
	}
	return int64(f[0])
}
