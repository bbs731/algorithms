package DFS

/***

给你一个二维字符网格数组 grid ，大小为 m x n ，你需要检查 grid 中是否存在 相同值 形成的环。

一个环是一条开始和结束于同一个格子的长度 大于等于 4 的路径。对于一个给定的格子，你可以移动到它上、下、左、右四个方向相邻的格子之一，可以移动的前提是这两个格子有 相同的值 。

同时，你也不能回到上一次移动时所在的格子。比方说，环  (1, 1) -> (1, 2) -> (1, 1) 是不合法的，因为从 (1, 2) 移动到 (1, 1) 回到了上一次移动时的格子。

如果 grid 中有相同值形成的环，请你返回 true ，否则返回 false 。
 */

func containsCycle(grid [][]byte) bool {
	m := len(grid)
	n := len(grid[0])

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	type pair struct {
		x, y int
	}
	// 从 1254 得到的教训， 这里写成 "global" 的变量， 然后下面去 reset 更靠谱，比 dfs 返回值靠谱，那个容易出错，虽然我还没明白为啥会错。
	l := 0
	found := false
	pm := make(map[pair]int)
	var dfs func(int, int, int, int, byte)
	dfs = func(x, y int, px, py int, target byte) {
		if visited[x][y] {
			return
		}

		visited[x][y] = true
		pm[pair{x, y}] = l
		l++

		dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
		for _, d := range dirs {
			r, c := x+d[0], y+d[1]
			if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] != target || (r == px && c == py) { // need additional parent check
				continue
			}

			if p, ok := pm[pair{r, c}]; ok {
				// we found a cycle
				if l-p >= 4 {
					found = true
				}
			}
			dfs(r, c, x, y, target)
		}
	}

	for i := range grid {
		for j := range grid[0] {
			if visited[i][j] == false {
				//reset
				found = false
				l = 0
				pm = make(map[pair]int)
				dfs(i, j, -1, -1, grid[i][j])
				if found {
					return true
				}
			}
		}
	}

	return false
}
