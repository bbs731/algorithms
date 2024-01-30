package DFS

/***

这道题，太难了， 噩梦啊！

这道题，点到命门了。
 */

func closedIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var res bool // 用global 的变量， 如果把这个写在  dfs 的返回值里， 就会出错。 原因是 visited[x][y] = true 的时候？ dfs 返回什么值？（你还需要判断一下 x, y 是不是 border) 看下面的例子吧，再写一遍。
	var dfs func(int, int)
	dfs = func(x, y int) {
		if visited[x][y] {
			return
		}
		dirs := [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
		visited[x][y] = true
		for _, d := range dirs {
			r, c := x+d[0], y+d[1]
			if r < 0 || r >= m || c < 0 || c >= n {
				res = false // reach border so mark this connected components group whole as false
				continue
			}
			if grid[r][c] == 0 {
				dfs(r, c)
			}
		}
		return
	}

	ans := 0
	for i := range grid {
		for j := range grid[0] {
			if visited[i][j] == false && grid[i][j] == 0 {
				res = true
				dfs(i, j)
				if res {
					ans++
				}
			}
		}
	}
	return ans
}
