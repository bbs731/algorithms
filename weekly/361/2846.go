package weekly

/*


现有一棵由 n 个节点组成的无向树，节点按从 0 到 n - 1 编号。给你一个整数 n 和一个长度为 n - 1 的二维整数数组 edges ，其中 edges[i] = [ui, vi, wi] 表示树中存在一条位于节点 ui 和节点 vi 之间、权重为 wi 的边。

另给你一个长度为 m 的二维整数数组 queries ，其中 queries[i] = [ai, bi] 。对于每条查询，请你找出使从 ai 到 bi 路径上每条边的权重相等所需的 最小操作次数 。在一次操作中，你可以选择树上的任意一条边，并将其权重更改为任意值。

注意：

查询之间 相互独立 的，这意味着每条新的查询时，树都会回到 初始状态 。
从 ai 到 bi的路径是一个由 不同 节点组成的序列，从节点 ai 开始，到节点 bi 结束，且序列中相邻的两个节点在树中共享一条边。
返回一个长度为 m 的数组 answer ，其中 answer[i] 是第 i 条查询的答案。

提示：

1 <= n <= 104
edges.length == n - 1
edges[i].length == 3
0 <= ui, vi < n
1 <= wi <= 26
生成的输入满足 edges 表示一棵有效的树
1 <= queries.length == m <= 2 * 104
queries[i].length == 2
0 <= ai, bi < n

 */

const MAXN = 1 << 15

type edge struct {
	to, wt int
}

func minOperationsQueries(n int, edges [][]int, queries [][]int) []int {

	var (
		// 为什么静态数组，会对 leetcode 不友好呢？
		//Log2   [MAXN]int
		//fa     [MAXN][15]int
		//weight [MAXN][15][26]int
		//dep    [MAXN]int
		//g      [][]edge

		Log2   []int
		fa     [][15]int
		weight [][15][26]int
		dep    []int
		g      [][]edge
	)

	Log2 = make([]int, n)
	fa = make([][15]int, n)
	weight = make([][15][26]int, n)
	dep = make([]int, n)

	var dfs func(int, int, int)

	dfs = func(cur, fath, d int) {
		dep[cur] = d
		fa[cur][0] = fath

		// 这个可以写在外层吗？ 可以的。  两种写法都可以， 1. 倍增写在 dfs 里面。  2. 倍增写在 dfs 之后。看下面的代码。
		//for i := 1; i <= Log2[dep[cur]]; i++ {
		//	fa[cur][i] = fa[fa[cur][i-1]][i-1]
		//	for j := 0; j < 26; j++ {
		//		weight[cur][i][j] = weight[cur][i-1][j] + weight[fa[cur][i-1]][i-1][j]
		//	}
		//}
		for _, e := range g[cur] {
			if w := e.to; w != fath {
				//weight[w][0] = e.wt
				weight[w][0][e.wt] = 1
				dfs(w, cur, d+1)
			}
		}
	}

	var lca func(int, int) (int, [26]int)
	lca = func(a, b int) (int, [26]int) {
		cnt := [26]int{}
		if dep[a] > dep[b] {
			a, b = b, a
		}

		for dep[a] != dep[b] {
			for j := 0; j < 26; j++ {
				cnt[j] += weight[b][Log2[dep[b]-dep[a]]][j]
			}
			b = fa[b][Log2[dep[b]-dep[a]]]
		}

		if a == b {
			return a, cnt
		}

		for k := Log2[dep[a]]; k >= 0; k-- {
			if fa[a][k] != fa[b][k] {
				for j := 0; j < 26; j++ {
					cnt[j] += weight[a][k][j] + weight[b][k][j]
				}
				a, b = fa[a][k], fa[b][k]
			}
		}
		for j := 0; j < 26; j++ {
			cnt[j] += weight[a][0][j] + weight[b][0][j]
		}
		return fa[a][0], cnt
	}

	for i := 2; i < n; i++ {
		Log2[i] = Log2[i/2] + 1
	}
	g = make([][]edge, n)
	// need to build graph
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]-1
		g[v] = append(g[v], edge{w, wt})
		g[w] = append(g[w], edge{v, wt})
	}

	dfs(0, -1, 0)

	for k := 1; k <= Log2[n-1]; k++ {
		for i := 0; i < n; i++ {
			if fa[i][k-1] != -1 {
				fa[i][k] = fa[fa[i][k-1]][k-1]
				for j := 0; j < 26; j++ {
					weight[i][k][j] = weight[i][k-1][j] + weight[fa[i][k-1]][k-1][j]
				}
			}
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		//cnt := [26]int{}
		p, cnt := lca(q[0], q[1])
		//fmt.Println(p, q[0], q[1], cnt)
		pathlen := dep[q[0]] + dep[q[1]] - 2*dep[p]
		maxCnt := 0
		for j := 0; j < 26; j++ {
			// 计算路径的过程，应该放到 lca 里面，虽然麻烦. 要不然，就需要在这里，再模拟一遍。
			maxCnt = max(maxCnt, cnt[j])
		}
		ans[i] = pathlen - maxCnt
	}

	return ans
}
