package BFS

/***
这是一道 BFS 经典的 template 好题啊
 */

func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	type pair struct {
		x, y int
	}

	q := []pair{}
	cnts := 0

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 2 {
				q = append(q, pair{i, j})
			} else if grid[i][j] == 1 {
				cnts++
			}
		}
	}
	// 特判
	if cnts == 0 {
		return 0
	}

	dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	// step 初始化成 1 的原因是，找到结果直接 return step, 不会执行 step++操作， 所以会初始化为1
	for step := 1; len(q) > 0; step++ {
		tmp := q
		q = nil

		for _, p := range tmp {
			x, y := p.x, p.y
			for _, d := range dirs {
				r, c := x+d[0], y+d[1]
				if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 || grid[r][c] == 2 {
					continue
				}
				if grid[r][c] == 1 {
					grid[r][c] = 2 // mark as rotten
					// 入队
					q = append(q, pair{r, c})
					cnts--
				}
			}
			if cnts == 0 {
				// got end
				return step
			}
		}
	}
	// not possible
	return -1
}
