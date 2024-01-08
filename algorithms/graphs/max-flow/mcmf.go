package max_flow

// 最小费用最大流 MCMF（即满流时的费用）
//https://oi-wiki.org/graph/flow/min-cost/

//下面的代码 minCostFlowSPFA 来自：
//https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/graph.go#L4308
//框架是 Edmonds-Karp 算法， 把 Edmonds-Karp 中发现增广路的 bfs() 改成 队列优化的Bellman-Ford 也就是 SPFA 算法。因为边权可能是负值，
//所以这里没有使用 Dijkstra 而使用 队列优化的 Belleman-Ford。
//mcmf 整个算法的时间复杂度是 O(nmf) (注意哈， Edmonds-Karp 计算最大流的算法的时间复杂度是 O（V*E^2) 的，没有depend on flow 的值域，
//但是基于 Edmonds-Karp 的 mcmf 的算法的时间复杂度是有f, flow 的值域项的）
//我们还知道 Dinic 再求最大流的时候时间复杂度是O（V^2*E) 如果用 Dinic 求 mcmf 时间复杂度还有 f, flow 的值域项吗？我猜测还是有的，如何证明？
//因为 Edmonds-Karp 改进了 Ford-Fulkerson 算法(Ford-Fulkerson 用 DFS去找增广路的话）让它不依赖 f 值域。 但是没办法让 mcmf 不依赖 f 的值域。 Dinic应该也不可以。
//另外 mcmf 是有快速算法的，网络单纯型法，（解决 linear programing 的 simplex 吗？） 但是太复杂了，留给以后吧。

// 模板题 https://www.luogu.com.cn/problem/P3381
// LC2850 建模 https://leetcode.cn/problems/minimum-moves-to-spread-stones-over-grid/
func (*graph) minCostFlowSPFA(n, st, end int, edges [][]int) (int, int) {
	const inf int = 1e18
	st--
	end--

	type neighbour struct{ to, rid, cap, cost, eid int } // rid 为反向边在邻接表中的下标。
	g := make([][]neighbour, n)

	addEdge := func(from, to, cap, cost, eid int) {
		g[from] = append(g[from], neighbour{to, len(g[to]), cap, cost, eid})
		g[to] = append(g[to], neighbour{from, len(g[from]) - 1, 0, -cost, -1}) // 无向图上 0 换成 cap
	}

	for i, e := range edges {
		v, w, edgeCap, edgeCost := e[0], e[1], e[2], e[3]
		addEdge(v, w, edgeCap, edgeCost, i)
	}

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

	return edmondsKarp()
}

// https://oi-wiki.org/graph/flow/min-cost/
// dinic 的算法，还是参考的灵神的代码. 灵神的dinic 代码，在计算 mcmf 的时候， dfs 会无限循环。
// https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/graph.go#L3938
// 灵神的这段，(找 maxflow 的模板题目，去验证灵神的这段 dinic 代码是否在 maxflow 问题上，可行）。 dinic 求 maxflow 的代码，没验证过，
// 直接用在 mcmf 上有bug 不行。 下面的 mcmfDinic 参照了
// https://oi-wiki.org/graph/flow/min-cost/ 这里 mcmf Dinic 的实现， 但是用的是灵神的 maxflow 的框架代码。

func (*graph) mcmfDinic(n, st, end int, edges [][]int) (int, int) {
	const inf int = 1e18
	st--
	end--

	type neighbour struct{ to, rid, cap, cost, eid int } // rid 为反向边在邻接表中的下标。
	g := make([][]neighbour, n)

	addEdge := func(from, to, cap, cost, eid int) {
		g[from] = append(g[from], neighbour{to, len(g[to]), cap, cost, eid})
		g[to] = append(g[to], neighbour{from, len(g[from]) - 1, 0, -cost, -1}) // 无向图上 0 换成 cap
	}

	for i, e := range edges {
		v, w, edgeCap, edgeCost := e[0], e[1], e[2], e[3]
		addEdge(v, w, edgeCap, edgeCost, i)
	}

	//上面是建图的代码

	//下面是 Dinic mcmf 的 template 代码
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

	minCost := 0
	iter := make([]int, len(g))
	var dfs func(int, int) int
	dfs = func(v int, minF int) int {
		if v == end {
			return minF
		}
		ans := 0
		inQ[v] = true
		for ; iter[v] < len(g[v]); iter[v]++ {
			e := &g[v][iter[v]]
			if w := e.to; !inQ[w] && e.cap > 0 && dist[w] == dist[v]+e.cost {
				if f := dfs(w, min(minF-ans, e.cap)); f > 0 {
					e.cap -= f
					g[w][e.rid].cap += f
					ans += f
					minCost += f * e.cost
					//return f
				}
			}
		}
		//return 0
		inQ[v] = false
		return ans
	}

	dinic := func() (maxFlow int) {
		for spfa() {
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

	return dinic(), minCost
}
