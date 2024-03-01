package dp

func validPartition(nums []int) bool {
	n := len(nums)
	f := make([]bool, n+1)
	f[n] = true

	for i := n - 2; i >= 0; i-- {
		if j := i + 1; j < n && nums[i] == nums[j] {
			f[i] = f[i] || f[i+2]
		}
		if j := i + 2; j < n && nums[i] == nums[i+1] && nums[i+1] == nums[j] {
			f[i] = f[i] || f[i+3]
		}
		if j := i + 2; j < n && nums[i]+1 == nums[i+1] && nums[i+1]+1 == nums[j] {
			f[i] = f[i] || f[i+3]
		}
	}
	return f[0]
}
