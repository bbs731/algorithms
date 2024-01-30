package DFS

/***

给你一个大小为 m x n 的二进制矩阵 grid 。

岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。

岛屿的面积是岛上值为 1 的单元格的数目。

计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。

 */

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	var dfs func(int, int) int
	dfs = func(i, j int) (cnts int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		cnts = 1
		dir := [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
		for _, d := range dir {
			cnts += dfs(i+d[0], j+d[1])
		}
		return cnts
	}

	ans := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				ans = max(ans, dfs(i, j))
			}
		}
	}
	return ans
}
