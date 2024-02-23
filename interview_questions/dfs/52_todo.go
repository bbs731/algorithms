package dfs

func totalNQueens(n int) int {
	ans := 0
	var dfs func(int, int)
	dplus := make([]int, 2*n)
	dminus := make([]int, 2*n)

	dfs = func(i int, mask int) {
		if i == n {
			ans += 1
			return
		}
		for j := 0; j < n; j++ {
			if mask&(1<<j) != 0 {
				continue
			}
			if dplus[i+j] != 0 || dminus[j-i+n] != 0 {
				continue
			}
			dplus[i+j] = 1
			dminus[j-i+n] = 1
			dfs(i+1, mask|(1<<j))
			dplus[i+j] = 0
			dminus[j-i+n] = 0
		}
	}
	dfs(0, 0)
	return ans
}
