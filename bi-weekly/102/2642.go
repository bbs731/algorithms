package weekly

const inf = int(1e9)

type Graph [][]int

func Constructor(n int, edges [][]int) Graph {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
			g[i][i] = 0
		}
	}

	for _, e := range edges {
		g[e[0]][e[1]] = e[2]
	}
	return g

}

func (this *Graph) AddEdge(edge []int) {
	g := *this
	g[edge[0]][edge[1]] = edge[2]
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	g := *this
	n := len(g)
	vis := make([]bool, n)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = inf
	}
	dist[node1] = 0

	// 朴树的 Dijkstra  复杂度是 O(n^2)
	for {
		// 找到当前最短路，去更新它的邻居的最短路，
		// 根据数学归纳法，dis[x] 一定是最短路长度
		x := -1
		for i := 0; i < n; i++ {
			if vis[i] {
				continue
			}
			if x < 0 || dist[i] < dist[x] {
				x = i
			}
		}
		// 没找到
		if x < 0 || dist[x] == inf {
			return -1
		}
		if x == node2 {
			return dist[x]
		}

		//找到了，更新邻居。
		vis[x] = true
		for y, w := range g[x] {
			if dist[x]+w < dist[y] {
				dist[y] = dist[x] + w
			}
		}
	}
}
