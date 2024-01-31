package BFS

/***

你现在手里有一份大小为 n x n 的 网格 grid，上面的每个 单元格 都用 0 和 1 标记好了。其中 0 代表海洋，1 代表陆地。

请你找出一个海洋单元格，这个海洋单元格到离它最近的陆地单元格的距离是最大的，并返回该距离。如果网格上只有陆地或者海洋，请返回 -1。

我们这里说的距离是「曼哈顿距离」（ Manhattan Distance）：(x0, y0) 和 (x1, y1) 这两个单元格之间的距离是 |x0 - x1| + |y0 - y1| 。

 */

func maxDistance(grid [][]int) int {
	n := len(grid)
	type pair struct {
		x, y, dist int
	}

	var q []pair
	for i := range grid {
		for j := range grid {
			if grid[i][j] == 1 {
				q = append(q, pair{i, j, 0})
			}
		}
	}

	ans := -1
	for len(q) > 0 {
		tmp := q
		q = nil

		for _, p := range tmp {
			x, y := p.x, p.y
			dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
			for _, d := range dirs {
				r, c := x+d[0], y+d[1]
				if r < 0 || r >= n || c < 0 || c >= n || grid[r][c] == 1 {
					continue
				}
				grid[r][c] = 1 // mark as visited
				q = append(q, pair{r, c, p.dist + 1})
				ans = max(ans, p.dist+1)
			}
		}
	}
	return ans
}
