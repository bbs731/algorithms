package dfs

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	n := len(candidates)

	var dfs func(int, int, []int)
	dfs = func(i int, sum int, l []int) {
		if i == n {
			if sum == target {
				ans = append(ans, append([]int{}, l...))
			}
			return
		}
		if sum > target {
			return
		}

		dfs(i+1, sum, l)
		if sum+candidates[i] <= target {
			l = append(l, candidates[i])
			dfs(i, sum+candidates[i], l)
			l = l[:len(l)-1]
		}
	}

	dfs(0, 0, []int{})
	return ans
}

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	n :=len(candidates)

	var dfs func(int, int, []int)
	dfs = func(i int, sum int, l []int) {
		if i == n {
			if sum == target {
				ans = append(ans, append([]int{}, l...))
			}
			return
		}
		dfs(i+1, sum, l)
		for k:=1; sum + candidates[i]*k <= target; k++ {
			for j:=1; j <=k; j++ {
				l = append(l, candidates[i])
			}
			dfs(i+1, sum+candidates[i]*k, l)
			l = l[:len(l)-k]
		}
	}

	dfs(0, 0, []int{})
	return ans
}

