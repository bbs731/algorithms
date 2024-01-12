package one_day_exercise

import "fmt"

/*

给你一个大小为 m x n ，由若干 0 和 1 组成的二维网格 grid ，其中 1 表示陆地， 0 表示水。岛屿 由水平方向或竖直方向上相邻的 1 （陆地）连接形成。

如果 恰好只有一座岛屿 ，则认为陆地是 连通的 ；否则，陆地就是 分离的 。

一天内，可以将 任何单个 陆地单元（1）更改为水单元（0）。

返回使陆地分离的最少天数。



示例 1：


输入：grid = [[0,1,1,0],[0,1,1,0],[0,0,0,0]]
输出：2
解释：至少需要 2 天才能得到分离的陆地。
将陆地 grid[1][1] 和 grid[0][2] 更改为水，得到两个分离的岛屿。
示例 2：


输入：grid = [[1,1]]
输出：2
解释：如果网格中都是水，也认为是分离的 ([[1,1]] -> [[0,0]])，0 岛屿。

 */

func minDays(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	nodeCount := 0
	labels := make(map[int]int)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				labels[i*n+j] = nodeCount
				nodeCount++
			}
		}
	}
	//edges := make([][]int, 0)
	g := make([][]int, nodeCount)

	dir := [][]int{[]int{0, -1}, []int{0, 1}, []int{-1, 0}, []int{1, 0}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				for k := 0; k < len(dir); k++ {
					ni := i + dir[k][0]
					nj := j + dir[k][1]
					if ni < 0 || ni >= m || nj < 0 || nj >= n {
						continue
					}
					if grid[ni][nj] == 1 {
						//edges = append(edges, []int{labels[i*n+j], labels[ni*n+nj]})
						g[labels[i*n+j]] = append(g[labels[i*n+j]], labels[ni*n+nj])
						g[labels[ni*n+j]] = append(g[labels[ni*n+nj]], labels[i*n+j])
					}
				}
			}
		}
	}

	cuts, scc := findCutVertices(nodeCount, g)

	fmt.Println(nodeCount, scc, cuts)
	if nodeCount == 0 || scc > 1 {
		return 0
	}

	// scc = 1
	if nodeCount == 1 || len(cuts) > 0 {
		return 1
	}

	return 2
}

// low(v): 在不经过 v 父亲的前提下能到达的最小的时间戳
func findCutVertices(n int, g [][]int) (cuts []int, scc int) {
	isCut := make([]bool, n)
	dfn := make([]int, n) // DFS 到终点的时间（从 1开始）
	dfsClock := 0
	low := make([]int, n) // low[v]  用它来存储不经过其父亲能到达的最小的时间戳

	var tarjan func(u, father int)
	tarjan = func(u, father int) {
		dfsClock++
		low[u] = dfsClock
		dfn[u] = dfsClock
		child := 0

		for _, v := range g[u] {
			if dfn[v] == 0 {
				child++
				tarjan(v, u)
				low[u] = min(low[u], low[v])

				if low[v] >= dfn[u] { // 以 v 为根的子树中没有反向边能连回 u 的祖先（或者连到u上， 也算割顶）
					isCut[u] = true
				}
			} else if v != father {
				low[u] = min(low[u], dfn[v])
			}
		}

		// 起点的逻辑，特殊处理。 特判：只有一个儿子的树根，删除后并没有增加连通分量的个数，这种情况下不是割顶
		if father == -1 && child == 1 {
			isCut[u] = false
		}
	}

	for u, timestamp := range dfn {
		if timestamp == 0 {
			scc++
			tarjan(u, -1)
		}
	}

	for u, is := range isCut {
		if is {
			cuts = append(cuts, u)
		}
	}
	return cuts, scc
}
