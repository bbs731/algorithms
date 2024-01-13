package one_day_exercise

import (
	"fmt"
	"slices"
	"sort"
)

// 没做出来， 下面的解答是错的，
// 先别把时间浪费在这门难得题上， 等等， 官方的解答。

type neighbour struct{ to, eid int }
func findBridges(n int, edges [][]int) (isBridge []bool) {
	g := make([][]neighbour, n)

	for i, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], neighbour{v, i})
		g[v] = append(g[v], neighbour{u, i})
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

func minimumCost(cost []int, roads [][]int) int64 {
	n := len(cost)
	g := make([][]neighbour,n)
	type pair struct {cnt, i int}
	deg := make([]pair, n)
	const inf int = 1e17
	for i, e := range roads {
		v, w := e[0], e[1]
		deg[v].cnt++
		deg[v].i = v
		deg[w].cnt++
		deg[w].i = w
		g[v] = append(g[v], neighbour{w, i})
		g[w] = append(g[w], neighbour{v, i})
	}

	bridge := findBridges(n, roads)
	c := make([]int, n)
	dcc := 0
	var dfs func(int)

	dfs = func(x int) {
		c[x]= dcc
		for _, y := range g[x]{
			if c[y.to] !=0 || bridge[y.eid]{
				continue
			}
			dfs(y.to)
		}
	}

	for i:=0; i<n; i++ {
		if c[i] == 0 {
			dcc++
			dfs(i)
		}
	}
	ans := 0


	// 特判   dcc == n  // 这是一条链子。
	if dcc == n || dcc == 1 {
		head := inf
		sort.Slice(deg, func(i, j int) bool { return deg[i].cnt < deg[j].cnt})
		ll := deg[0].cnt
		for i:=0; i<n; i++ {
			if deg[i].cnt == ll {
				head = min(head, cost[i])
			} else {
				break
			}
		}
		return int64(head)
	}

	// dcc  从 1 ,2,... dcc
	dcc_cost := make([]int, dcc)
	for i:=0; i<dcc; i++ {
		dcc_cost[i] = inf
	}

	for i:=0; i<n; i++ {
		dcc_cost[c[i]-1] = min(dcc_cost[c[i]-1], cost[i])
	}
	sort.Ints(dcc_cost)
	for i:=0; i<dcc-1; i++{
		ans += dcc_cost[i]
	}
	return int64(ans)
}