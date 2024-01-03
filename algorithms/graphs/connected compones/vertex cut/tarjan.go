package vertex_cut

//代码来自：
//https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/graph.go#L742

type graph struct{}

func (*graph) findCutVertices(n int, g [][]int) (isCut []bool) {
	isCut = make([]bool, n)
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

				if low[v] >= dfn[v] { // 以 v 为根的子树中没有反向边能连回 u 的祖先（或者连到u上， 也算割顶）
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
			tarjan(u, -1)
		}
	}

	cuts := []int{}
	for u, is := range isCut {
		if is {
			cuts = append(cuts, u)
		}
	}
	return
}
