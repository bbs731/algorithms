package DFS

/***

给你一个下标从 0 开始大小为 m x n 的二维整数数组 grid ，其中下标在 (r, c) 处的整数表示：

如果 grid[r][c] = 0 ，那么它是一块 陆地 。
如果 grid[r][c] > 0 ，那么它是一块 水域 ，且包含 grid[r][c] 条鱼。
一位渔夫可以从任意 水域 格子 (r, c) 出发，然后执行以下操作任意次：

捕捞格子 (r, c) 处所有的鱼，或者
移动到相邻的 水域 格子。
请你返回渔夫最优策略下， 最多 可以捕捞多少条鱼。如果没有水域格子，请你返回 0 。

格子 (r, c) 相邻 的格子为 (r, c + 1) ，(r, c - 1) ，(r + 1, c) 和 (r - 1, c) ，前提是相邻格子在网格图内。

 */

func findMaxFish(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	var dfs func(int, int) int
	dfs = func(x, y int) (cnts int) {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
			return
		}
		cnts = grid[x][y]
		grid[x][y] = 0 // 省去了 visited 数组
		dir := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, d := range dir {
			cnts += dfs(x+d[0], y+d[1])
		}
		return cnts
	}

	ans := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] != 0 {  // 这个判断，都可以省了
				ans = max(ans, dfs(i, j))
			}
		}
	}
	return ans
}

/***
https://leetcode.cn/problems/maximum-number-of-fish-in-a-grid/solutions/2250953/wang-ge-tu-dfs-mo-ban-pythonjavacgo-by-e-lykw/
灵神的题解： 一模一样， 赞！
 */
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findMaxFish(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
			return 0
		}
		sum := grid[x][y]
		grid[x][y] = 0 // 标记成访问过
		for _, d := range dirs { // 四方向移动
			sum += dfs(x+d.x, y+d.y)
		}
		return sum
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return
}
