package chagne_root_DP

/*
 834 的变形题目。 换根DP的典型题目。
灵神的题解：
https://leetcode.cn/problems/minimum-edge-reversals-so-every-node-is-reachable/solutions/2445681/mo-ban-huan-gen-dppythonjavacgojs-by-end-8qiu/

关键在于： ans[y] = ans[x] +  direction(x->y)

时间复杂度两遍 DFS： O(n)

 */
func minEdgeReversals(n int, edges [][]int) []int {
	g := make([][][2]int, n)

	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], [2]int{y, 1})
		g[y] = append(g[y], [2]int{x, -1})
	}

	var dfs func(int, int)
	ans := make([]int, n)
	dfs = func(x, fa int) {
		for _, p := range g[x] {
			if p[0] != fa {
				if p[1] == -1 {
					ans[0] += 1 // adding cost
				}
				dfs(p[0], x)
			}
		}
	}
	dfs(0, -1)

	var reroot func(int, int)
	reroot = func(x, fa int) {
		for _, p := range g[x] {
			y := p[0]
			if y != fa {
				ans[y] = ans[x] + p[1]
				reroot(y, x)
			}
		}
	}
	reroot(0, -1)
	return ans
}
