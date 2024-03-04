package Bipartite_Graph

func possibleBipartition(n int, dislikes [][]int) bool {
	g := make([][]int, n)

	for _, d := range dislikes {
		v, w := d[0]-1, d[1]-1 // index 是从 1 开始算的额，所以可以减去1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	colors := make([]int, n)

	var dfs func(int, int) bool
	dfs = func(v int, c int) bool {
		colors[v] = c

		for _, w := range g[v] {
			if colors[w] == c || colors[w] == 0 && !dfs(w, 3^c) {
				return false
			}
		}
		return true
	}

	for i, c := range colors {
		if c != 0 {
			continue
		}

		if !dfs(i, 1) {
			return false
		}
	}
	return true
}
