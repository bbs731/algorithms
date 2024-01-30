package DFS

/***

给定一个 row x col 的二维网格地图 grid ，其中：grid[i][j] = 1 表示陆地， grid[i][j] = 0 表示水域。

网格中的格子 水平和垂直 方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。

岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。

 */

/****

这道题，给我们不一样的启示，就是，边界检查的位置， 可以写在：
1. dfs 的开头。
2. 调用 dfs 之前。
 */

func islandPerimeter(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(int, int) int
	dfs = func(x, y int) (cnts int) {
		if visited[x][y] {
			return
		}
		visited[x][y] = true
		dir := [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

		for _, d := range dir {
			r, c := x+d[0], y+d[1]
			if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 { // 这道题，说明什么？ 边界检查可以写在 dfs 的开头， 也可以写在 call dfs 之前。
				cnts += 1 // 如果周边是海水 cnt + 1
				continue
			}
			cnts += dfs(x+d[0], y+d[1])
		}
		return
	}

	ans := 0
	for i := range grid {
		for j := range grid[0] {
			if visited[i][j] == false && grid[i][j] == 1 {
				ans += dfs(i, j)
			}
		}
	}
	return ans
}
