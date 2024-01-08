package topological_sort

import "slices"

func (*graph) topoSort(n int, edges [][]int) []int {
	g := make([][]int, n)
	visit := make([]int, n) // 0 unvisited, -1 visiting,  1 visited
	orders := make([]int, 0, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w) // Topological sort 适用于 DAG， 有向图， 所以不要加 w 到 v 的边
	}

	var dfs func(int) bool
	dfs = func(u int) bool {
		visit[u] = -1
		for _, v := range g[u] {
			if visit[v] == -1 {
				return false // 这是返祖边， 代表DAG 有环。
			}
			if visit[v] == 0 && !dfs(v) {
				return false
			}
		}
		visit[u] = 1
		orders = append(orders, u)
		return true
	}

	for i := 0; i < n; i++ {
		if visit[i] == 0 && !dfs(i) {
			// report false, cycle detected in DAG
			return nil
		}
	}

	slices.Reverse(orders)
	return orders
}
