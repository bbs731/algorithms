package BFS

type pair struct {
	x, y, dist int
}

func bfs_grid_template(grid [][]int) {

	m := len(grid)
	n := len(grid[0])

	// 有些题目，可以原地修改 grid 省去 visited 数组
	visited := make([][]bool, m)

	// 初始化 Q
	q := []pair{}

	for len(q) > 0 {
		// 这里是技巧， BFS 一层一层遍历。
		tmp := q
		q = nil
		for _, p := range tmp {
			x, y := p.x, p.y
			dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
			for _, d := range dirs {
				r, c := x+d[0], y+d[1]
				if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 {
					continue
				}
				grid[r][c] = 0 // mark as visited
				q = append(q, pair{r, c, p.dist + 1})
				//res[r][c] = p.dist + 1
			}
		}
	}

}
