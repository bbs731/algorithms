package BFS

type pair struct {
	x, y int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	if grid[0][0] == 1 {
		return -1
	}
	if m == 1 {
		return 1
	}
	// 我的天啊， 写对一道题，真是太难了。
	// 如果  grid 只有一个元素， 那么应该返回 1 但是， 下面的解会返回 -1 这是一个 corner case
	q := []pair{{}}

	for step := 1; len(q) > 0; step++ {
		tmp := q
		q = nil
		dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {-1, -1}, {-1, 1}, {1, 1}, {1, -1}}
		for _, p := range tmp {
			x, y := p.x, p.y

			for _, d := range dirs {
				nx, ny := x+d[0], y+d[1]
				if nx < 0 || nx >= m || ny < 0 || ny >= n || grid[nx][ny] == 1 {
					continue
				}

				if nx == m-1 && ny == n-1 {
					return step + 1
				}
				// mark as visited
				grid[nx][ny] = 1
				q = append(q, pair{nx, ny})
			}
		}
	}
	return -1
}
