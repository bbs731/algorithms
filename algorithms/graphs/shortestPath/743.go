package shortestPath

import "container/heap"

// 标标准准的 Dijstra Shortest 模版题
const inf int = 1e18
func networkDelayTime(times [][]int, n int, k int) int {
	// 这里唯一的难点，就是 node index 是 从 1 开始，不是从 0 开始的。 哈哈！
	d := shortestPathDijkstra(n+1, k, times)

	ans := 0
	for i:=1; i<=n; i++ {
		x := d[i]
		if x == inf {
			return -1
		}
		ans = max(ans, x)
	}
	return ans
}

func shortestPathDijkstra(n, st int, edges [][]int) (dist []int) {
	type neighbor struct{ to, wt int }
	g := make([][]neighbor, n)
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]
		g[v] = append(g[v], neighbor{w, wt})
		//g[w] = append(g[w], neighbor{v, wt})
	}

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
