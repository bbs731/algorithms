package DFS

/***
不要迷信， 你可以的。
 */

// 1 是陆地， 0 是海洋
func numEnclaves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	var border bool
	var cnts int
	var dfs func(int, int)
	dfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			border = true
			return
		}
		if grid[x][y] != 1 {
			return
		}

		grid[x][y] = 0 // mark to 0, 这样省去 visited 数组
		cnts++

		dirs := [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} // 考， dirs 里面写出了，重复元素， 哎！
		for _, d := range dirs {
			dfs(x+d[0], y+d[1])
		}
	}

	ans := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				cnts = 0
				border = false
				dfs(i, j)
				if !border {
					ans += cnts
				}
			}
		}
	}
	return ans
}

func numEnclaves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	// 这题可以省略 visited 数组， 直接在 grid 上把 1 改成 0
	//visited := make([][]bool, m)
	//for i := range visited {
	//	visited[i] = make([]bool, n)
	//}

	var border bool
	var cnts int
	var dfs func(int, int)
	dfs = func(x, y int) {
		if grid[x][y] != 1 {
			return
		}
		//if visited[x][y] {
		//	return
		//}

		grid[x][y] = 0 // mark to 0, 这样省去 visited 数组
		cnts++
		//visited[x][y] = true

		dirs := [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} // 我靠 dirs 写错了，有重复的元素， 没copy就出错？
		for _, d := range dirs {
			r, c := x+d[0], y+d[1]
			if r < 0 || r >= m || c < 0 || c >= n {
				border = true
				continue
			}
			dfs(r, c)
		}
	}

	ans := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				cnts = 0
				border = false
				dfs(i, j)
				if !border {
					ans += cnts
				}
			}
		}
	}
	return ans
}
