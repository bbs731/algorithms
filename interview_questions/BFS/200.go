package BFS

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	ans := 0

	var dfs func(int, int)
	dfs = func(i, j int) {
		// mark as visited
		grid[i][j] = byte('2')

		for _, d := range dirs {
			x, y := i+d[0], j+d[1]
			if x >= m || x < 0 || y >= n || y < 0 {
				continue
			}

			if grid[x][y] == '1' {
				dfs(x, y)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				ans++
			}
		}
	}
	return ans
}
