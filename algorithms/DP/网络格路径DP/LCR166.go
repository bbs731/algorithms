package dp

/***
	f[i][j] = max(f[i-1][j], f[i][j-1]) + frame[i][j]
	f[i+1][j+1] = max(f[i][j+1], f[i+1][j]) + frame[i][j]
 */

func jewelleryValue(frame [][]int) int {
	m := len(frame)
	n := len(frame[0])

	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			f[i+1][j+1] = max(f[i][j+1], f[i+1][j]) + frame[i][j]
		}
	}
	return f[m][n]
}
