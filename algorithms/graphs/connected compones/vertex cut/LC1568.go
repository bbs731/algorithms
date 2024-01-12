package vertex_cut

import "fmt"

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
					// 这里容易错
					v := labels[i*n+j]
					w := labels[ni*n+nj]
					if grid[ni][nj] == 1 {
						g[v] = append(g[v], w)
						g[w] = append(g[w], v)
					}
				}
			}
		}
	}

	cuts, scc := findCutVertices(nodeCount, g)

	fmt.Println(nodeCount, scc, len(scc), cuts)
	if nodeCount == 0 || len(scc) > 1 {
		return 0
	}

	// scc = 1
	if nodeCount == 1 || len(cuts) > 0 {
		return 1
	}

	return 2
}

// 下面这个 找割点的，写的无比的强大， 不但返回割点，还能返回 scc components, 本来 Tarjan 就能做这两件事。
// 以前的举例，都是要么单独找 SCC 要们，单独计算 cut vertex or cut edges.  没有结合过。
// 下面的写法，结和了 scc + cut vertex, 主要改的点是 else 判断  inSt 而不是  v != father

// 本题，本来只需要简单返回 scc 的计数就可以， 炫技了！
// low(v): 在不经过 v 父亲的前提下能到达的最小的时间戳
func findCutVertices(n int, g [][]int) (cuts []int, scc [][]int) {
	scc = [][]int{}
	isCut := make([]bool, n)
	dfn := make([]int, n) // DFS 到终点的时间（从 1开始）
	dfsClock := 0
	low := make([]int, n) // low[v]  用它来存储不经过其父亲能到达的最小的时间戳
	inSt := make([]bool, n)
	st := []int{}

	var tarjan func(u, father int)
	tarjan = func(u, father int) {
		dfsClock++
		low[u] = dfsClock
		dfn[u] = dfsClock
		child := 0

		st = append(st, u)
		inSt[u] = true
		for _, v := range g[u] {
			if dfn[v] == 0 {
				child++
				tarjan(v, u)
				low[u] = min(low[u], low[v])

				if low[v] >= dfn[u] { // 以 v 为根的子树中没有反向边能连回 u 的祖先（或者连到u上， 也算割顶）
					isCut[u] = true
				}
			} else if inSt[v] { // 这里要十分的注意， 不能用  v!=father
				//} else if v != father {
				low[u] = min(low[u], dfn[v])
			}
		}

		// 起点的逻辑，特殊处理。 特判：只有一个儿子的树根，删除后并没有增加连通分量的个数，这种情况下不是割顶
		if father == -1 && child == 1 {
			isCut[u] = false
		}
		if dfn[u] == low[u] {
			comp := []int{}
			for {
				v := st[len(st)-1]
				st = st[:len(st)-1]
				inSt[v] = false
				comp = append(comp, v)
				if v == u { // v is one of the SCC root. See the video explaination
					break
				}
			}
			scc = append(scc, comp)
		}
	}

	for u, timestamp := range dfn {
		if timestamp == 0 {
			//scc++
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
