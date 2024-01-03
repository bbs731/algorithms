package edge_cut

// https://oi-wiki.org/graph/cut/#_

type graph struct{}

func (*graph) findBridges(n int, edges [][]int) (isBridge []bool) {
	type neightbour struct{ to, eid int }
	g := make([][]neightbour, n)

	for i, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], neightbour{v, i})
		g[v] = append(g[v], neightbour{u, i})
	}

	isBridge = make([]bool, len(edges))
	dfn := make([]int, len(g)) // 值从 1 开始
	low := make([]int, len(g))
	dfsClock := 0

	var tarjan func(int, int)
	tarjan = func(u, fid int) { // 是用 fid 跟父节点的 edge id, 可以兼容重边的情况
		dfsClock++
		dfn[u] = dfsClock
		low[u] = dfsClock

		for _, e := range g[u] {
			if v := e.to; dfn[v] == 0 {
				tarjan(v, e.eid)
				low[u] = min(low[u], low[v])
				if low[v] > dfn[u] { //以 v 为根的子树中没有反向边能连回 u 或 u 的祖先，所以 u-v 必定是桥
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

	bridgeEIDs := []int{}
	for eid, b := range isBridge {
		if b {
			bridgeEIDs = append(bridgeEIDs, eid)
		}
	}
	return
}
