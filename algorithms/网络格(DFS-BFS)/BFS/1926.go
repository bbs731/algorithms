package BFS

/***

https://leetcode.cn/problems/nearest-exit-from-entrance-in-maze/solutions/869437/go-bfs-by-endlesscheng-k2cu/
灵神的题解， 学习一下，用作 BFS 的模板。

 */

type pair struct{ x, y int }

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func nearestExit(g [][]byte, entrance []int) int {
	n, m := len(g), len(g[0])
	s := pair{entrance[0], entrance[1]}
	g[s.x][s.y] = 0 // 重置成 0， 省去了 visited 数组。
	q := []pair{s}
	for ans := 1; len(q) > 0; ans++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			for _, d := range dir4 {
				if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < n && 0 <= y && y < m && g[x][y] == '.' {
					if x == 0 || y == 0 || x == n-1 || y == m-1 {
						return ans
					}
					g[x][y] = 0
					q = append(q, pair{x, y})
				}
			}
		}
	}
	return -1
}

func nearestExit(maze [][]byte, entrance []int) int {
	m := len(maze)
	n := len(maze[0])

	q := [][2]int{{entrance[0], entrance[1]}}
	maze[entrance[0]][entrance[1]] = 0 // 原地修改，可以省去 visited 数组。
	for step := 1; len(q) > 0; step++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
			for _, d := range dirs {
				r, c := p[0]+d[0], p[1]+d[1]
				if r >= 0 && r < m && c >= 0 && c < n && maze[r][c] == '.' {
					if r == 0 || r == m-1 || c == 0 || c == n-1 {
						return step
					}
					q = append(q, [2]int{r, c})
					maze[r][c] = 0
					//if p[0] != entrance[0] || p[1] != entrance[1] {
					//	return step
					//}
					//continue
				}
			}
		}
	}
	return -1
}

func nearestExit(maze [][]byte, entrance []int) int {
	m := len(maze)
	n := len(maze[0])

	q := [][2]int{{entrance[0], entrance[1]}}
	maze[entrance[0]][entrance[1]] = 0 // 原地修改，可以省略 visited 数组。
	for step := 0; len(q) > 0; step++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
			for _, d := range dirs {
				r, c := p[0]+d[0], p[1]+d[1]
				if r < 0 || r >= m || c < 0 || c >= n {
					if p[0] != entrance[0] || p[1] != entrance[1] {
						return step
					}
					continue
				}
				// otherwise add to the queue
				if maze[r][c] == '.' {
					q = append(q, [2]int{r, c})
					maze[r][c] = 0
				}
			}
		}
	}
	return -1
}
