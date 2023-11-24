package backtracking

func permute(nums []int) [][]int {
	ans := [][]int{}
	path := []int{}
	n := len(nums)

	var dfs func(int, []bool)

	dfs = func(i int, set []bool) {
		if i == n {
			// copy path, since path is global variable
			ans = append(ans, append([]int(nil), path...))
			return
		}

		for j := 0; j < n; j++ {
			if set[j] == false {
				set[j] = true
				path = append(path, nums[j])
				dfs(i+1, set)
				path = path[:len(path)-1]
				set[j] = false
			}
		}
	}

	set := make([]bool, n)
	dfs(0, set)
	return ans
}
