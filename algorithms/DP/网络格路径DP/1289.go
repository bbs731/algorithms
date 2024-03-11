package dp

import "math"

func minFallingPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 {
				f[i][j] = 0
			} else {
				f[i][j] = math.MaxInt32
			}
		}
	}

	/***
		f[i][j] = min(f[i-1][j+1], f[i-1][j-1]) + grid[i][j] 理解错误
		f[i][j] = min (f[i-1][k], for all k that diff than j) + grid[i][j]
	*/

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				// i =0, j=0, k=0 的时候， 需要更新一下 f[1][1] 当 m =n =1 的时候。
				if k == j && i != 0 {
					continue
				}
				f[i+1][j+1] = min(f[i+1][j+1], f[i][k+1]+grid[i][j])
			}
		}
	}

	ans := math.MaxInt32
	for i := 1; i <= n; i++ {
		ans = min(ans, f[m][i])
	}
	return ans
}
