package weekly

/*
前后缀分解 : 模版题是 238
*/

func constructProductMatrix(grid [][]int) [][]int {
	n, m := len(grid), len(grid[0])
	mod := 12345

	p := make([][]int, n)
	for i := 0; i < n; i++ {
		p[i] = make([]int, m)
		//for j := 0; j < m; j++ {
		//	p[i][j] = 1    第一遍直接用 suffix 来 赋值的话，就省去了 初始化 1
		//}
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			p[i][j] = suffix
			suffix = (suffix * grid[i][j]) % mod
		}
	}
	prefix := 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			p[i][j] = (p[i][j] *prefix) %mod
			prefix = (prefix * grid[i][j]) % mod
		}
	}

	return p
}
