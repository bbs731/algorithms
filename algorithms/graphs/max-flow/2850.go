package max_flow

func minimumMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	st := m * n                                          // 超级源点
	end := st + 1                                        // 超级汇点
	type neighbour struct{ to, rid, cap, cost, eid int } // rid 为反向边在邻接表中的下标。
	g := make([][]neighbour, m*n+2)

	addEdge := func(from, to, cap, cost, eid int) {
		g[from] = append(g[from], neighbour{to, len(g[to]), cap, cost, eid})
		g[to] = append(g[to], neighbour{from, len(g[from]) - 1, 0, -cost, -1}) // 无向图上 0 换成 cap
	}

	// 难点，就在于建图了！
	for x, row := range grid {
		for y, v := range row {
			if v > 1 {
				addEdge(st, x*n+y, v-1, 0, x)
				for i, r := range grid {
					for j, w := range r {
						if w == 0 {
							addEdge(x*n+y, i*n+j, 1, abs(x-i)+abs(y-j), i)
						}
					}
				}
			} else if v == 0 {
				addEdge(x*n+y, end, 1, 0, y)
			}
		}
	}

	// 下面是最小费用最大流模板
	const inf int = 1e9
	dist := make([]int, len(g))
	type vi struct{ v, i int }
	fa := make([]vi, len(g))
	inQ := make([]bool, len(g))
	spfa := func() bool {
		for i := range dist {
			dist[i] = inf
		}
		dist[st] = 0
		inQ[st] = true
		q := []int{st}

		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			inQ[v] = false

			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				if newD := dist[v] + e.cost; newD < dist[w] {
					dist[w] = dist[v] + e.cost
					fa[w] = vi{v, i}
					if !inQ[w] {
						inQ[w] = true
						q = append(q, w)
					}
				}
			}
		}
		return dist[end] != inf
	}
	edmondsKarp := func() (maxFlow, minCost int) {
		for spfa() {
			// 沿着 spfa 找到的最短路，来增广
			minF := inf

			for v := end; v != st; {
				p := fa[v]
				if c := g[p.v][p.i].cap; c < minF {
					minF = c
				}
				v = p.v
			}
			for v := end; v != st; {
				p := fa[v]
				e := &g[p.v][p.i]
				e.cap -= minF
				g[v][e.rid].cap += minF // 这里太奇妙了， 如何建的图？
				v = p.v
			}
			maxFlow += minF
			minCost += dist[end] * minF
		}
		return
	}

	_, cost := edmondsKarp()
	return cost
}

func abs(x int) int {
	if x < 0 {
		return -x
	};
	return x
}
