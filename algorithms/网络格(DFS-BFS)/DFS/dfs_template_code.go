package DFS

/***

下面是 dfs grid template code
思考方式分为两种 ：
1. 正向， 下面的代码就是示例。
2. 反向， 如果正向考虑有难度，可以从反向的，边界出发解决问题。 譬如题目： 1034 和 417

 */
func grid_dfs_template(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	target := 2
	// 根据题目的需要， 如果可以 原地修改 grid的值不会乱，那么可以省略 visited 数组。
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var res bool // 建议使用 global variable 返回 dfs 里的解，建议不要用 dfs 的返回值，容易错, 除非特别方便。
	var dfs func(int, int)

	dfs = func(x, y int) {
		if visited[x][y] {
			return
		}
		visited[x][y] = true
		grid[x][y] = -1 // 如果可以修改 grid 的值，可以替代 visited 数组。
		dirs := [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

		for _, d := range dirs {
			r, c := x+d[0], y+d[1]
			if r < 0 || r >= m || c < 0 || c >= n {
				res = false // reach border so mark this connected components group whole as false
				continue
			}
			if grid[r][c] == target {
				dfs(r, c)
			}
		}
	}

	// loop matrix to find answer
	ans := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == target {
				res = true
				dfs(i, j)
				if res {
					ans++
				}
			}
		}
	}
	return ans
}
