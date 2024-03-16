package dfs

/***
简单的， dfs 遍历，记录 path
 */
func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	ans := [][]int{}
	var dfs func(int, []int)
	dfs = func(i int, p []int) {
		if i == n-1 {
			// record ans
			p = append(p, n-1)
			ans = append(ans, append([]int{}, p...))
			return
		}
		p = append(p, i)
		for _, j := range graph[i] {
			dfs(j, p)
		}
	}
	dfs(0, []int{})
	return ans
}
