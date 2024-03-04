package topological_sort

// kahn's algorithm to get topological sort order
//code from:
//https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/graph.go#L3105

type graph struct{}

func (*graph) topoSort(n int, edges [][]int) []int {
	g := make([][]int, n)
	deg := make([]int, len(g))

	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		deg[w]++
	}

	//fa := make([]int, len(g))
	//for i := range fa {
	//	fa[i] = -1
	//}
	//levels := make([]int, len(g))

	q := make([]int, 0, len(g))
	orders := q
	l := make([]int, 0, len(g))

	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
			// levels[i] = 1
			// NOTE: 若起点有特殊性，可以在这里初始化 dp
		}
	}

	for len(q) > 0 {
		v := q[0]
		l = append(l, v)
		q = q[1:]

		for _, w := range g[v] {
			// dp[w] = max(dp[w], dp[v]
			if deg[w]--; deg[w] == 0 {
				//fa[w] = v
				//levels[w] = levels[v] +1
				q = append(q, w)
			}
		}
	}

	// NOTE: 若 cap(q) 大于 0 则说明图中有环
	// 这个看不懂啊！ 看207 模板题目的例子吧。用一个 list 保存 toplogical sort 的结果。
	orders = orders[:len(g)-cap(q)]

	// 如果 l == len(g) == n 代表存在拓扑排序， 否则代表有环啊！
	l == n
	// NOTE: 若要重复求拓扑排序记得拷贝一份 deg

	{
		fa := make([]int, len(g))

		// EXTRA: path from end to start
		var end = len(g) - 1
		path := make([]int, 0, len(g))
		for v := end; v != -1; v = fa[v] {
			path = append(path, v)
		}
	}

	return orders
}
