package dfs

import "sort"

func getAncestors(n int, edges [][]int) [][]int {
	parent := make([]map[int]struct{}, n)
	deg := make([]int, n)
	g := make([][]int, n)

	for i := 0; i < n; i++ {
		parent[i] = make(map[int]struct{})
	}

	for _, e := range edges {
		from, to := e[0], e[1]
		deg[to]++
		g[from] = append(g[from], to)
		//g[to] = append(g[to], from)
	}

	q := []int{}
	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		u := q[len(q)-1]
		q = q[:len(q)-1]
		for _, v := range g[u] {
			deg[v]--
			pv := parent[v]
			pv[u] = struct{}{}
			for k := range parent[u] {
				pv[k] = struct{}{}
			}
			if deg[v] == 0 {
				q = append(q, v)
			}
		}
	}

	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		for k := range parent[i] {
			ans[i] = append(ans[i], k)
		}
		sort.Ints(ans[i])
	}
	return ans
}
