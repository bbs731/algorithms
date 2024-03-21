package dp

import (
	"math"
)

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	m := len(triangle[n-1])

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			f[i][j] = math.MaxInt32
		}
	}
	// 这个初始化， 成了关键！
	f[0][0] = 0

	ans := math.MaxInt32

	for i := 0; i < n; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			f[i+1][j+1] = min(f[i][j+1], f[i][j]) + triangle[i][j]
		}
	}

	for j := 0; j < len(triangle[n-1]); j++ {
		ans = min(ans, f[n][j+1])
	}

	return ans
}
