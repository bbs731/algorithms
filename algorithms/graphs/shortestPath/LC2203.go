package shortestPath

// 灵神的题解：
//https: //leetcode.cn/problems/minimum-weighted-subgraph-with-the-required-paths/solutions/1332967/by-endlesscheng-2mxm/
// 天才班的想法， 枚举中间，交汇的点！

const inf int = 1e18

type neighbor struct{ to, wt int }

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	rg := make([][]neighbor, n)
	g := make([][]neighbor, n)

	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]
		g[v] = append(g[v], neighbor{w, wt})
		rg[w] = append(rg[w], neighbor{v, wt})
	}

	d1 := shortestPathDijkstra(n, src1, g)
	d2 := shortestPathDijkstra(n, src2, g)
	d3 := shortestPathDijkstra(n, dest, rg)

	ans := inf
	for i := 0; i < n; i++ {
		ans = min(ans, d1[i]+d2[i]+d3[i])
	}

	if ans >= inf {
		return -1
	}
	return int64(ans)
}

func shortestPathDijkstra(n, st int, g [][]neighbor) (dist []int) {

	dist = make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[st] = 0
	from := make([]int, n)
	for i := range from {
		from[i] = -1
	}
	h := dijkstraHeap{{st, 0}}
	for len(h) > 0 {
		p := h.pop()
		v := p.v
		// 下面循环中的 newD < dist[w] 可能会把重复的节点 w 入堆
		// 也就是说，堆中可能会包含多个相同节点，且这些相同节点的 dist 值互不相同
		// 那么这个节点第二次及其后面出堆的时候，由于 dist[v] 已经更新成最短路了，可以直接跳过
		if p.dis > dist[v] {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newD := dist[v] + e.wt; newD < dist[w] {
				dist[w] = newD
				from[w] = v
				h.push(dijkstraPair{w, dist[w]})
			}
		}
	}
	return
}

type dijkstraPair struct{ v, dis int }
type dijkstraHeap []dijkstraPair

func (h dijkstraHeap) Len() int             { return len(h) }
func (h dijkstraHeap) Less(i, j int) bool   { return h[i].dis < h[j].dis }
func (h dijkstraHeap) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *dijkstraHeap) Push(v any)          { *h = append(*h, v.(dijkstraPair)) }
func (h *dijkstraHeap) Pop() (v any)        { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *dijkstraHeap) push(v dijkstraPair) { heap.Push(h, v) }
func (h *dijkstraHeap) pop() dijkstraPair   { return heap.Pop(h).(dijkstraPair) }
