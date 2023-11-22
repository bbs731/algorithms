package backtracking

// LC 78
func subsets(nums []int) [][]int {
	var ans [][]int
	var path []int
	var dfs func(int)

	n := len(nums)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]int(nil), path...))
			return // 第一遍写， 忘了这里的 return 大哥啊，长点心啊！
		}

		// 不选 nums[i]
		dfs(i + 1)

		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1]
	}

	dfs(0)
	return ans
}
