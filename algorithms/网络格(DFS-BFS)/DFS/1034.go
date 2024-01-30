package DFS

/***
给你一个大小为 m x n 的整数矩阵 grid ，表示一个网格。另给你三个整数 row、col 和 color 。网格中的每个值表示该位置处的网格块的颜色。

如果两个方块在任意 4 个方向上相邻，则称它们 相邻 。

如果两个方块具有相同的颜色且相邻，它们则属于同一个 连通分量 。

连通分量的边界 是指连通分量中满足下述条件之一的所有网格块：

在上、下、左、右任意一个方向上与不属于同一连通分量的网格块相邻
在网格的边界上（第一行/列或最后一行/列）
请你使用指定颜色 color 为所有包含网格块 grid[row][col] 的 连通分量的边界 进行着色。

并返回最终的网格 grid 。



示例 1：

输入：grid = [[1,1],[1,2]], row = 0, col = 0, color = 3
输出：[[3,3],[3,2]]
示例 2：

输入：grid = [[1,2,2],[2,3,2]], row = 0, col = 1, color = 3
输出：[[1,3,3],[2,3,3]]
示例 3：

输入：grid = [[1,1,1],[1,1,1],[1,1,1]], row = 1, col = 1, color = 2
输出：[[2,2,2],[2,1,2],[2,2,2]]

 */

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m := len(grid)
	n := len(grid[0])

	// save the connected components
	com := make([][]bool, m)
	for i := range com {
		com[i] = make([]bool, n)
	}
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	var dfs func(int, int, int)
	dfs = func(x, y, z int) {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != z || com[x][y] { // 这里是个坑，没有com[x][y] 的判断，  dfs 不会结束。 哎！掉坑里了。
			return
		}
		com[x][y] = true
		for _, d := range dirs {
			dfs(x+d[0], y+d[1], z)
		}
	}
	// 染色
	dfs(row, col, grid[row][col])

	isBorder := func(x, y int) bool {
		for _, d := range dirs {
			r, c := x+d[0], y+d[1]
			if r < 0 || r >= m || c < 0 || c >= n || com[r][c] == false {
				return true
			}
		}
		return false
	}

	for i := range grid {
		for j := range grid[0] {
			if com[i][j] == true && isBorder(i, j) {
				grid[i][j] = color
			}
		}
	}
	return grid
}

/***
明显的，上面的代码是可以优化的。

赞， 就是变量的名字，容易弄乱了， c for color for column 哎！
 */

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m := len(grid)
	n := len(grid[0])

	// save the connected components
	com := make([][]int, m) // 1 代表是在同一个 connected component, 另外，如果是 2 代表不但是同一个 connected components 还是 border
	for i := range com {
		com[i] = make([]int, n)
	}
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	var dfs func(int, int, int)
	dfs = func(x, y, z int) {
		if com[x][y] > 0 { // 这里是个坑，没有com[x][y] 的判断，  dfs 不会结束。 哎！掉坑里了。
			return
		}
		com[x][y] = 1
		for _, d := range dirs {
			r, c := x+d[0], y+d[1]
			if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] != z { // 根据题目的需要， 边界的检查也可以放在这里。
				com[x][y] = 2 // mark [x,y] as border
				continue
			}
			dfs(x+d[0], y+d[1], z)
		}
	}
	dfs(row, col, grid[row][col])

	for i := range grid {
		for j := range grid[0] {
			if com[i][j] == 2 { //2 stands for border
				grid[i][j] = color
			}
		}
	}
	return grid
}
