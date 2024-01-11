package max_flow

/*
1. 建图的思路是对的。
2. 如果要求，最大费，需要把 cost 取反。 最后返回的 minCost 需要取反。
3. 恭喜你，你的 dinic 实现，是对的，已经验证了两道题目了。 (dfs return 0, 还是 return ans 的两种都是对的）
 */
func maximumANDSum(nums []int, numSlots int) int {

	const inf int = 1e18
	n := len(nums)
	m := numSlots
	st := n + m
	end := n + m + 1

	type neighbor struct{ to, rid, cap, cost int } // rid 为反向边在邻接表中的下标。
	g := make([][]neighbor, n+m+2)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost})
	}
	for i, num := range nums {
		addEdge(st, i, 1, 0)
		for j := 1; j <= m; j++ {
			// 这里灵神说的，有点问题， cap 可以去 1
			addEdge(i, n+j-1, 1, -(num & j)) // 这里的 cost 要取负，因为要求，最大费。
		}
	}
	for i := 0; i < m; i++ {
		addEdge(n+i, end, 2, 0)
	}

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

		inQ[v] = true
		for ; iter[v] < len(g[v]); iter[v]++ {
			e := &g[v][iter[v]]
			if w := e.to; !inQ[w] && e.cap > 0 && dist[w] == dist[v]+e.cost {
				if f := dfs(w, min(minF, e.cap)); f > 0 {
					e.cap -= f
					g[w][e.rid].cap += f
					minCost += f * e.cost
					inQ[v] = false
					return f
				}
			}
		}
		//return 0
		inQ[v] = false
		return 0
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
	dinic()
	return -minCost
}
