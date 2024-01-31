package BFS

/***
给你一个大小为 n x n 的二元矩阵 grid ，其中 1 表示陆地，0 表示水域。

岛 是由四面相连的 1 形成的一个最大组，即不会与非组内的任何其他 1 相连。grid 中 恰好存在两座岛 。

你可以将任意数量的 0 变为 1 ，以使两座岛连接起来，变成 一座岛 。

返回必须翻转的 0 的最小数目。



示例 1：

输入：grid = [[0,1],[1,0]]
输出：1
示例 2：

输入：grid = [[0,1,0],[0,0,0],[0,0,1]]
输出：2

 */

/***
先用要找到，这连个 connected components , id 为 1， 和 2
用BFS， 一层，一层，把周边的 grid 从 0 mark 到 1 或者 2 直到 1 和 2相遇为止


这道题，很难一次写对啊！ 非常好的，题目啊。 就是面试的时候，写的太长了，还是容易出租哦。

问题: 如何写出一个双向的 BFS ？

 */

func shortestBridge(grid [][]int) int {
	n := len(grid)
	dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	var dfs func(int, int)
	dfs = func(x, y int) {
		if grid[x][y] == 2 {
			return
		}
		grid[x][y] = 2 // mark to 2 as visited
		for _, d := range dirs {
			r, c := x+d[0], y+d[1]
			if r < 0 || r >= n || c < 0 || c >= n || grid[r][c] != 1 {
				continue
			}
			dfs(r, c)
		}
	}

	var a, b int
	for i := range grid {
		for j := range grid {
			if grid[i][j] == 1 {
				a, b = i, j
				break // 这里是一个坑， break 不能跳出两层循环，所以不能在这里直接 call dfs, 应该只记录 a,b 然后在外面 call dfs
			}
		}
	}
	dfs(a, b)
	// now we have 2 SCC,  one marked all as 1 , the other marked as 2

	type pair struct {
		x, y, id int
	}
	q := []pair{}
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
	}

	// initialize q
	for i := range grid {
		for j := range grid {
			if grid[i][j] == 1 {
				q = append(q, pair{i, j, 1})
			} else if grid[i][j] == 2 {
				q = append(q, pair{i, j, 2})
			}
		}
	}

	ans := 2*n + 1
	for len(q) > 0 {
		tmp := q
		q = nil

		for _, p := range tmp {
			x, y := p.x, p.y
			for _, d := range dirs {
				r, c := x+d[0], y+d[1]
				if r < 0 || r >= n || c < 0 || c >= n || grid[r][c] == p.id {
					continue
				}
				if grid[r][c] == 0 {
					grid[r][c] = p.id // mark [r,c] as visited
					dist[r][c] = dist[x][y] + 1
					q = append(q, pair{r, c, p.id})
					continue
				}
				// otherwise we find the meet, then can return the answer
				ans = min(ans, dist[r][c]+dist[x][y]) // 坑2 ： 这里不能直接 return , 有个test case 过不去，应该取最小值，结束时候返回。。
			}
		}
	}
	return ans
}
