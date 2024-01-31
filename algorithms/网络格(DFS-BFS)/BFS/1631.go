package BFS

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func minimumEffortPath(heights [][]int) int {
	inf := int(1e7)
	m := len(heights)
	n := len(heights[0])

	costs := make([][]int, m)
	for i := range costs {
		costs[i] = make([]int, n)
		for j := range costs[i] {
			costs[i][j] = inf
		}
	}

	dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	type pair struct {
		x, y int
	}
	q := []pair{{0, 0}}
	costs[0][0] = 0

	// 这实现的，就是 dijkstra 吗？ 不是，感觉更像 Bellman-Ford 最多松弛 |V|-1 次。总的复杂度是 |V||E|,  复杂度应该很高 (但是，可以用来检测负权）。
	// |V| = n^2   |E| = 2n   最后复杂度 O(n^3) 对于本题。 如果是 Dijkstra 的话， O（|E| + |V|)*log|V| 的复杂度，使用 heap 的话, 会快很多啊。
	for len(q) > 0 {
		tmp := q
		q = nil

		for _, p := range tmp {
			x, y := p.x, p.y
			for _, d := range dirs {
				r, c := x+d[0], y+d[1]
				if r < 0 || r >= m || c < 0 || c >= n {
					continue
				}
				// update r, c if can, then requeue   // 不明白为什么是对的， 不知道如何证明， 但是遇到两道题 (1293)，都可以这样来解决。
				if costs[r][c] > max(costs[x][y], abs(heights[r][c], heights[x][y])) {
					costs[r][c] = max(costs[x][y], abs(heights[r][c], heights[x][y]))
					q = append(q, pair{r, c})
				}
			}
		}
	}
	return costs[m-1][n-1]
}
