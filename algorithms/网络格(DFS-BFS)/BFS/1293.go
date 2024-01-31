package BFS

/***

给你一个 m * n 的网格，其中每个单元格不是 0（空）就是 1（障碍物）。每一步，您都可以在空白单元格中上、下、左、右移动。
如果您 最多 可以消除 k 个障碍物，请找出从左上角 (0, 0) 到右下角 (m-1, n-1) 的最短路径，并返回通过该路径所需的步数。如果找不到这样的路径，则返回 -1 。



这个题型没有见过啊。
根据灵神之前 BFS 的模板写的，用了技巧，（重新入队）题目是过了，但是不知道用的对还是不对， 不能给出证明。

 */

func shortestPath(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])

	type pair struct {
		x, y int
	}

	dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	visited := make(map[pair]int)
	start := pair{0, 0}
	visited[start] = 0
	q := []pair{start}

	// [[0]] 需要特判一下这个例子。
	if start.x == m-1 && start.y == n-1 {
		return 0
	}

	for step := 1; len(q) > 0; step++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			x, y := p.x, p.y
			for _, d := range dirs {
				r, c := x+d[0], y+d[1]
				if r < 0 || r >= m || c < 0 || c >= n {
					continue
				}
				if v, ok := visited[pair{r, c}]; !ok {
					// first time visited
					q = append(q, pair{r, c})
					visited[pair{r, c}] = visited[p] + grid[r][c] // grid[r][c] 是 0 就 + 0 是 1 就加 1 个cost
				} else {
					// update if can. 如果不更新可以吗？ 是否 BFS 满足，第一次visited 给的就是最短距离？
					// 对于这道题来说，不行，必须要用可能的更小值，更新。 如果产生了更新，还必须重新入队。
					// 哎，这里完全是靠感觉，来懵的， 理论依据是什么呢？
					visited[pair{r, c}] = min(v, visited[p]+grid[r][c])
					if v > visited[p]+grid[r][c] {
						q = append(q, pair{r, c})
					}
				}
				// 需要在这里检查解，而不能放在循环结束. 因为解依赖 step 的值，而不是 visited 保证的值。
				// 有的 test case 依赖这个
				if r == m-1 && c == n-1 {
					if visited[pair{r, c}] <= k {
						return step
					}
				}
			}
		}
	}

	return -1
}
