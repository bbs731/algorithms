package BFS

/***

你还记得那条风靡全球的贪吃蛇吗？

我们在一个 n*n 的网格上构建了新的迷宫地图，蛇的长度为 2，也就是说它会占去两个单元格。蛇会从左上角（(0, 0) 和 (0, 1)）开始移动。我们用 0 表示空单元格，用 1 表示障碍物。蛇需要移动到迷宫的右下角（(n-1, n-2) 和 (n-1, n-1)）。

每次移动，蛇可以这样走：

如果没有障碍，则向右移动一个单元格。并仍然保持身体的水平／竖直状态。
如果没有障碍，则向下移动一个单元格。并仍然保持身体的水平／竖直状态。

如果它处于水平状态并且其下面的两个单元都是空的，就顺时针旋转 90 度。蛇从（(r, c)、(r, c+1)）移动到 （(r, c)、(r+1, c)）。


如果它处于竖直状态并且其右面的两个单元都是空的，就逆时针旋转 90 度。蛇从（(r, c)、(r+1, c)）移动到（(r, c)、(r, c+1)）。

返回蛇抵达目的地所需的最少移动次数。
如果无法到达目的地，请返回 -1。


造作：
1. 向右。
2. 向下
3. 旋转。
每个造作都有两种情况需要考虑， 水平和垂直的情况。


太爽了， 一次救过！  还是要学习一下，灵神的题解和模板
https://leetcode.cn/problems/minimum-moves-to-reach-target-with-rotations/solutions/2093126/huan-zai-if-elseyi-ge-xun-huan-chu-li-li-tw8b/

 */

type pair struct {
	x, y    int
	horizon bool // horizontal for vertical?
	// if horizon occupy [x,y] and [x, y+1]
	// if vertical occupy [x,y] and [x+1, y]
}

func minimumMoves(grid [][]int) int {
	n := len(grid)
	visited := make(map[pair]int) // int keep the shortest dist can reached from start point
	start := pair{0, 0, true}
	q := []pair{start}
	visited[start] = 0

	for len(q) > 0 {
		tmp := q
		q = nil
		for _, p := range tmp {
			candidates := []pair{}
			if p.horizon {
				// move right
				if p.y+1+1 < n && grid[p.x][p.y+1+1] == 0 {
					//next = pair{p.x, p.y + 1, p.horizon}
					candidates = append(candidates, pair{p.x, p.y + 1, p.horizon})

				}
				// move down  and rotate 可以公用一个条件。
				if p.x+1 < n && grid[p.x+1][p.y] == 0 && grid[p.x+1][p.y+1] == 0 {
					//next = pair{p.x + 1, p.y, p.horizon}
					candidates = append(candidates, pair{p.x + 1, p.y, p.horizon})
					candidates = append(candidates, pair{p.x, p.y, false})

				}

			} else {
				// vertical
				// move right
				if p.y+1 < n && grid[p.x][p.y+1] == 0 && grid[p.x+1][p.y+1] == 0 {
					candidates = append(candidates, pair{p.x, p.y + 1, false}) // move right
					candidates = append(candidates, pair{p.x, p.y, true})      // rotate up
				}
				// move down
				if p.x+1+1 < n && grid[p.x+1+1][p.y] == 0 {
					candidates = append(candidates, pair{p.x + 1, p.y, false}) // move down
				}
			}

			for _, next := range candidates {
				if v, ok := visited[next]; ok {
					visited[next] = min(v, visited[p]+1)
					// ToDo: ? 这里还需要入队吗？ < v 的情况下？test case 有不过的情况下，就考虑一下， ruguo  < v 则入队，否则不入队
					// 这里不需要重新列队了， 在网络格（grid) 里， 的BFS ，就是有求最短路的性质。
				} else {
					//first time visited
					visited[next] = visited[p] + 1
					q = append(q, next)
				}
			}
		}
	}

	end := pair{n - 1, n - 2, true}
	if visited[end] == 0 {
		return -1
	}
	return visited[end]
}

/***
尝试一下 模板， 让 visited 数组， 只保存 true or false
下面的代码也是， 可以的。
 */
func minimumMoves(grid [][]int) int {
	n := len(grid)
	visited := make(map[pair]bool) // int keep the shortest dist can reached from start point
	start := pair{0, 0, true}
	q := []pair{start}
	visited[start] = true

	for step := 1; len(q) > 0; step++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			candidates := []pair{}
			if p.horizon {
				// move right
				if p.y+1+1 < n && grid[p.x][p.y+1+1] == 0 {
					//next = pair{p.x, p.y + 1, p.horizon}
					candidates = append(candidates, pair{p.x, p.y + 1, p.horizon})

				}
				// move down  and rotate 可以公用一个条件。
				if p.x+1 < n && grid[p.x+1][p.y] == 0 && grid[p.x+1][p.y+1] == 0 {
					//next = pair{p.x + 1, p.y, p.horizon}
					candidates = append(candidates, pair{p.x + 1, p.y, p.horizon})
					candidates = append(candidates, pair{p.x, p.y, false})

				}

			} else {
				// vertical
				// move right
				if p.y+1 < n && grid[p.x][p.y+1] == 0 && grid[p.x+1][p.y+1] == 0 {
					candidates = append(candidates, pair{p.x, p.y + 1, false}) // move right
					candidates = append(candidates, pair{p.x, p.y, true})      // rotate up
				}
				// move down
				if p.x+1+1 < n && grid[p.x+1+1][p.y] == 0 {
					candidates = append(candidates, pair{p.x + 1, p.y, false}) // move down
				}
			}

			for _, next := range candidates {
				if _, ok := visited[next]; !ok {
					//first time visited
					visited[next] = true
					q = append(q, next)
					if next.x == n-1 && next.y == n-2 {
						return step
					}
				}
			}
		}
	}
	return -1
}
