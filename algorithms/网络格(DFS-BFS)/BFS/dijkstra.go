package BFS

import "container/heap"

type dijkstraPair struct {
	v, dis int
}
type dijkstraHeap []dijkstraPair

// 这是大堆还是小堆啊？
func (h dijkstraHeap) Len() int           { return len(h) }
func (h dijkstraHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h dijkstraHeap) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h *dijkstraHeap) Push(v any)         { *h = append(*h, v.(dijkstraPair)) }
func (h *dijkstraHeap) Pop() (v any) {
	a := *h
	*h, v = a[:len(a)-1], a[:len(a)-1]
	return
}

// golang 自己写一个最小堆好难啊！
func (h *dijkstraHeap) pop() dijkstraPair {
	return heap.Pop(h).(dijkstraPair)
}
func (h *dijkstraHeap) push(v dijkstraPair) {
	//*h = append(*h, v)
	heap.Push(h, v)
}

type neighbour struct {
	to, wt int
}

func shortestPathDijkstra(m, n, st int, grid [][]int) (dist []int) {
	// grid[x, y] 的 node 编号，  x*m + j
	g := make([][]neighbour, m*n)

	dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for i := range grid {
		for j := range grid[0] {
			for _, d := range dirs {
				x, y := i+d[0], j+d[1]
				if x < 0 || x >= m || y < 0 || y >= n {
					continue
				}
				g[i*m+j] = append(g[i*m+j], neighbour{x*m + y, grid[x][y]})
				g[x*m+y] = append(g[x*m+y], neighbour{i*m + j, grid[i][j]})
			}
		}
	}
	const inf int = 1e18
	dist = make([]int, m*n)
	for i := range dist {
		dist[i] = inf
	}
	// st = 0  // 初始化
	dist[st] = 0
	h := dijkstraHeap{{st, 0}}

	for len(h) > 0 {
		p := h.pop()
		v := p.v

		if p.dis > dist[v] {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newD := dist[v] + e.wt; newD < dist[w] {
				dist[w] = newD
				h.push(dijkstraPair{w, dist[w]})
			}
		}
	}
	return
}
