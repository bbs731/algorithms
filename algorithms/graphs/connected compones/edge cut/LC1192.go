package edge_cut

func criticalConnections(n int, connections [][]int) [][]int {
	type neightbour struct{ to, eid int }

	g := make([][]neightbour, n)
	for i, e := range connections {
		u, v := e[0], e[1]
		g[u] = append(g[u], neightbour{v, i})
		g[v] = append(g[v], neightbour{u, i})
	}

	dfsClock := 0
	dfn := make([]int, n)
	low := make([]int, n)
	isBridge := make([]bool, len(connections))

	var tarjan func(int, int)
	tarjan = func(u, fid int) {
		dfsClock++
		dfn[u] = dfsClock
		low[u] = dfsClock

		for _, e := range g[u] {
			v := e.to
			if dfn[v] == 0 {
				tarjan(v, e.eid)
				low[u] = min(low[u], low[v])
				if low[v] > dfn[u] { // 这里是最关键的代码 比较的是 low[v] > dfn[u]
					isBridge[e.eid] = true
				}
			} else if e.eid != fid {
				low[u] = min(low[u], dfn[v])
			}
		}
	}

	for u, timestamp := range dfn {
		if timestamp == 0 {
			tarjan(u, -1)
		}
	}

	ans := make([][]int, 0)
	for i, b := range isBridge {
		if b {
			ans = append(ans, connections[i])
		}
	}
	return ans
}
