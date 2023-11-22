package backtracking

// 第一遍就写错了， 再练
// 从 answer 的视角 (选哪个数）
func subsets(nums []int) [][]int {

	var ans [][]int
	var path []int
	var dfs func(int)

	n := len(nums)
	dfs = func(i int) {
		a := make([]int, len(path))
		copy(a, path)
		ans = append(ans, a)

		// 技巧1： ans = append(ans, append([]int(nil), path...)) // copy path

		if i == n {
			return
		}

		for j := i; j < n; j++ {
			path = append(path, nums[j])
			dfs(j + 1)
			l := len(path)
			path = path[:l-1]

			// 技巧 2 path = path[:len(path)-1]
		}
	}

	dfs(0)
	return ans
}
