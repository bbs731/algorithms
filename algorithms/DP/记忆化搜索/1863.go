package dp

func subsetXORSum(nums []int) int {
	n := len(nums)

	ans := 0
	var dfs func(int, int)
	dfs = func(i, mask int) {
		if i == n {
			ans += mask
			return
		}
		dfs(i+1, mask)
		dfs(i+1, mask^nums[i])
	}

	dfs(0, 0)
	return ans
}
