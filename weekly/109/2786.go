package weekly

func maxScore(nums []int, x int) int64 {
	n := len(nums)
	cache := make([][2]int, n)
	for i := range cache {
		cache[i][0] = -1
		cache[i][1] = -1
	}

	var dfs func (int, int) int
	dfs = func(i int, j int) int {
		if i == n {
			return 0
		}
		if cache[i][j] != - 1{
			return cache[i][j]
		}
		// 不选
		res := dfs(i+1, j)
		//选
		if nums[i]%2 == j {
			res = max(res, dfs(i+1,j) + nums[i])
		}else {
			res = max(res, dfs(i+1, nums[i]%2) + nums[i] - x )
		}
		cache[i][j]= res
		return res
	}
	// 这道题的，题目要求， 0 必须选。
	return int64(dfs(1, nums[0]%2) + nums[0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
