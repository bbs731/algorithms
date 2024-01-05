package max_flow

type graph struct{}

func (*graph) maxFlowDinic(n, st, end int, edges [][]int) int {
	const inf int = 1e18
	st--
	end--

	// 还是这样建图直观啊！ 用向前星建图，总是得反应一下
	type neightbour struct{ to, rid, cap, eid int } // rid 为反向边在领接表的下标。 eid >0 是正向边， <0 为反向边
	g := make([][]neightbour, n)

	addEdge := func(from, to, cap, eid int) {
		g[from] = append(g[from], neightbour{to, len(g[to]), cap, eid})
		g[to] = append(g[to], neightbour{from, len(g[from]) - 1, 0, -1}) // 这里 eid 标记为-1, 表示反边。
	}
	for i, e := range edges {
		v, w, edgeCap := e[0], e[1], e[2]
		addEdge(v, w, edgeCap, i)
	}

	d := make([]int, len(g)) // d[i] 数组表示 vertex i 离开 st 的距离， 越大，代表离 end 越近。
	bfs := func() bool {
		clear(d)
		d[st] = 1
		q := []int{st}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, e := range g[v] {
				if w := e.to; e.cap > 0 && d[w] == 0 {
					d[w] = d[v] + 1
					q = append(q, w)
				}
			}
		}
		return d[end] > 0
	}

	// 寻找增广路
	iter := make([]int, len(g)) // 当前弧，在其之前的边已经没有用了，避免对没有用的边进行多次检查
	var dfs func(int, int) int
	dfs = func(v, minF int) int {
		if v == end {
			return minF
		}
		for ; iter[v] < len(g[v]); iter[v]++ {
			e := &g[v][iter[v]]
			if w := e.to; e.cap > 0 && d[w] > d[v] {
				if f := dfs(w, min(minF, e.cap)); f > 0 {
					e.cap -= f
					g[w][e.rid].cap += f
					return f
				}
			}
		}
		return 0
	}

	dinic := func() (maxFlow int) {
		for bfs() {
			clear(iter)
			for {
				if f := dfs(st, inf); f > 0 {
					maxFlow += f
				} else {
					break
				}
			}
		}
		return
	}

	maxFlow := dinic()

	// EXTRA: 容量复原
	for _, es := range g {
		for i, e := range es {
			if e.eid > 0 { // 正向边
				es[i].cap += g[e.to][e.rid].cap // 这里用 es[i] 而不用 e 因为 e 是个 copy, 我们的更改不应该生效在副本上。
				g[e.to][e.rid] = 0
			}
		}
	}
	return maxFlow
}
