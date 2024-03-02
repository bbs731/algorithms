package dfs

func reachableNodes(n int, edges [][]int, restricted []int) int {
	g := make([][]int, n)
	for _, e := range edges {
		from, to := e[0], e[1]
		g[from]= append(g[from], to)
		g[to]= append(g[to], from)
	}
	rm := make(map[int]bool, n)
	for _, r := range restricted {
		rm[r]= true
	}

	var dfs func(int, int) int
	dfs = func(r int, p int) int {
		tot := 0
		for _, x := range g[r] {
			if !rm[x]&& x!=p {
				tot += dfs(x, r)
			}
		}
		return tot + 1
	}
	return dfs(0, -1)
}
