package dp

import "math"

// DP 的解法， 初始化是难点
func minPathSum(grid [][]int) int {
	inf := math.MaxInt32
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = inf
		}
	}
	dp[0][1] = 0
	dp[1][0] = 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i+1][j+1] = min(dp[i][j+1], dp[i+1][j]) + grid[i][j]
		}
	}

	return dp[m][n]
}

/***
可以用 0-1 BFS 的算法，来求最短距离。
当然，也可用 Dijstra
 */
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dis := make([][]int, m)
	inf := math.MaxInt32

	type pair struct{ x, y int }
	var dir4 = []struct{ x, y int }{{1, 0}, {0, 1}}

	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = inf
		}
	}
	dis[0][0] = grid[0][0]
	q := []pair{{0, 0}} // 两个 slice 头对头来实现 deque
	for len(q) > 0 {
		var p pair
		p, q = q[0], q[1:]
		for _, d := range dir4 {
			x, y := p.x+d.x, p.y+d.y
			if 0 <= x && x < m && 0 <= y && y < n {
				g := grid[x][y]
				if dis[p.x][p.y]+g < dis[x][y] {
					dis[x][y] = dis[p.x][p.y] + g
					q = append(q, pair{x, y})
				}
			}
		}
	}
	return dis[m-1][n-1]
}

/****
这个写法，太奇特了，
dfs 不加入 cache 时， 会超时。  加入了 cache 不是为了得到 dist 的数值， 而是为了剪枝。
第一次看到用 cache 来剪枝， 自己是不是已经走火入魔了？
 */

func minPathSum(grid [][]int) int {
	ans := math.MaxInt32
	m := len(grid)
	n := len(grid[0])
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = -1
		}
	}

	var dfs func(int, int, int)
	dfs = func(i, j int, sum int) {
		if i == m-1 && j == n-1 {
			ans = min(ans, sum)
			return
		}
		if sum >= ans {
			return
		}

		if cache[i][j] != -1 && sum >= cache[i][j] {
			return
		}

		cache[i][j] = sum

		dirs := [][2]int{{0, 1}, {1, 0}}
		for _, d := range dirs {
			x, y := i+d[0], j+d[1]
			if x >= m || y >= n {
				continue
			}
			dfs(x, y, sum+grid[x][y])
		}
	}

	dfs(0, 0, grid[0][0])
	return ans
}


/***
改成 缓存值的形式， 好理解。 dfs（i，j) 定义的是 从  (i, j) 到 （m-1, n-1) 的最小距离。
那下面的代码，就好理解了。
 */
func minPathSum(grid [][]int) int {
	inf := math.MaxInt32
	m := len(grid)
	n := len(grid[0])
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = -1
		}
	}
	cache[m-1][n-1] = grid[m-1][n-1]

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i >= m && j >= n {
			return inf
		}

		if cache[i][j] != -1 {
			return cache[i][j]
		}

		ans := inf
		dirs := [][2]int{{0, 1}, {1, 0}}
		for _, d := range dirs {
			x, y := i+d[0], j+d[1]
			if x >= m || y >= n {
				continue
			}
			ans = min(ans, dfs(x, y)+grid[i][j])
		}
		cache[i][j] = ans
		return ans
	}

	dfs(0, 0)
	return cache[0][0]
}
