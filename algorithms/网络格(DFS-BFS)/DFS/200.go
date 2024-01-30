package DFS

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0' // mark 成 visited
		dirs := [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
		for _, d := range dirs {
			dfs(i+d[0], j+d[1])
		}
		return
	}

	ans := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				dfs(i, j)
				ans++
			}
		}
	}
	return ans
}
