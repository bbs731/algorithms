package BFS

import "container/heap"

/****
LC2290
Dijkstra 版本的实现，完全验证了，之前分析的结论。
1. 自己最初实现的方法 （Bellman_Ford), 元素可以重新入队， 时间复杂度 O(VE)  测试集上的耗时  1740ms
2. 转换成最短路问题，都是正边，时间复杂度 O((V+E)*logV)  耗时 602ms
3. 灵神给的 0-1 BFS 用 deque 入队，而且保证每个 grid node 入队一次，时间复杂度是线性的 O(V) 节点个数时间复杂度，耗时 197ms

 */

func minimumObstacles(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	d := shortestPathDijkstra(m, n, 0, grid)
	return d[(m-1)*n+n-1]
}

type dijkstraPair struct {
	v, dis int
}
type dijkstraHeap []dijkstraPair

func (h dijkstraHeap) Len() int      { return len(h) }
func (h dijkstraHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// 这是大堆还是小堆啊？  这里决定了是最小堆。
func (h dijkstraHeap) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h *dijkstraHeap) Push(v any)        { *h = append(*h, v.(dijkstraPair)) }

// heap 已经 internally  把heap 最小的元素，置换到了，数组的末尾，等待我们去移除。
func (h *dijkstraHeap) Pop() (v any) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

//golang 自己写一个最小堆好难啊！
func (h *dijkstraHeap) push(v dijkstraPair) { heap.Push(h, v) }
func (h *dijkstraHeap) pop() dijkstraPair   { return heap.Pop(h).(dijkstraPair) }

type neighbour struct {
	to, wt int
}

func shortestPathDijkstra(m, n, st int, grid [][]int) (dist []int) {
	// grid[x, y] 的 node 编号，  x*m + j
	g := make([][]neighbour, m*n)

	dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	// 建图
	for i := range grid {
		for j := range grid[0] {
			for _, d := range dirs {
				x, y := i+d[0], j+d[1]
				if x < 0 || x >= m || y < 0 || y >= n {
					continue
				}
				// grid[i][j]  到  grid[x][y] 的边
				// node 的编号规则 ：  (i, j) = i*n + j
				g[i*n+j] = append(g[i*n+j], neighbour{x*n + y, grid[x][y]})
				g[x*n+y] = append(g[x*n+y], neighbour{i*n + j, grid[i][j]})
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
