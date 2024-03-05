package shortestPath

import "container/heap"

type graph struct{}

func (*graph) shortestPathDijkstra(n, st int, edges [][]int) (dist, cnts []int) {
	MOD := int(1e9) + 7
	type neighbor struct{ to, wt int }
	g := make([][]neighbor, n)
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]
		g[v] = append(g[v], neighbor{w, wt})
		g[w] = append(g[w], neighbor{v, wt})
	}

	const inf int = 1e18
	dist = make([]int, n)
	cnts = make([]int, n)
	for i := range dist {
		dist[i] = inf
		cnts[i] = 1
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
				cnts[w] = max(cnts[v], 1) // 这里需要 cnts[w] = cnts[v] 如果所有的 cnts 都初始化为 1 的话
				from[w] = v
				h.push(dijkstraPair{w, dist[w]})
			} else if newD == dist[w] {
				cnts[w] += cnts[v] // 这里是重点！
				cnts[w] %= MOD
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

func countPaths(n int, roads [][]int) int {
	g := &graph{}
	_, cnts := g.shortestPathDijkstra(n, 0, roads)
	return cnts[n-1]
}
