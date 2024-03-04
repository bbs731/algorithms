package Bipartite_Graph

/*

模板代码来自：
https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/graph.go

判断，是否是 bipartite graph
 */

func isBipartite(g [][]int) bool {

	colors := make([]int, len(g)) // n nodes
	var dfs func(int, int) bool
	dfs = func(v int, c int) bool {
		colors[v] = c
		for _, w := range g[v] {
			// 这个条件，写的太牛了。
			if colors[w] == c || colors[w] == 0 && !dfs(w, 3^c) {
				return false
			}
		}
		return true
	}

	for i, c := range colors {
		if c != 0 {
			// has been marked
			continue
		}
		if !dfs(i, 1) {
			return false
		}
	}
	return true
}
