package Strongly_Connected_Component

import "slices"

// code from  endlesswang  codeforces-go
// https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/graph.go#L3301
// video explanation:  https://www.youtube.com/watch?v=wUgWX0nc4NY&list=PLDV1Zeh2NRsDGO4--qE8yH72HFL1Km93P&index=23
// low[N]
// 定义： low-link value of a node is the smallest(lowest) node id reachable from that node when doing a DFS (including itself). 用它来存储不经过其父亲能到达的最小的时间戳

// 时间复杂度是 O(n+m) 线性的。
func sccTarjan(g [][]int) ([][]int, []int) {
	scc := [][]int{}
	dfn := make([]int, len(g))
	low := make([]int, len(g))
	dfsClock := 0
	st := []int{}
	inSt := make([]bool, len(g))
	var tarjan func(int)
	tarjan = func(u int) {
		dfsClock++
		dfn[u] = dfsClock
		low[u] = dfsClock // low-link

		st = append(st, u)
		inSt[u] = true
		for _, v := range g[u] {
			if dfn[v] == 0 { //unvisited
				tarjan(v)
				low[u] = min(low[u], low[v]) // u->v 这条边，是搜索树里的边。
			} else if inSt[v] {
				low[u] = min(low[u], dfn[v]) // u->v 这条边，不是搜索树里的边的情况。
			}

			//https://www.youtube.com/watch?v=hKhLj7bfDKk&list=PLDV1Zeh2NRsDGO4--qE8yH72HFL1Km93P&index=24
			//https://github.com/williamfiset/Algorithms/blob/master/src/main/java/com/williamfiset/algorithms/graphtheory/TarjanSccSolverAdjacencyList.java
			// 这样写，也是对的吧！
			//if dfn[v] == 0 {
			//	tarjan(v)
			//}
			//if inSt[v] {
			//	low[u] = min(low[u], low[v])
			//}
		}

		if dfn[u] == low[u] {
			comp := []int{}
			for {
				v := st[len(st)-1]
				st = st[:len(st)-1]
				inSt[v] = false
				comp = append(comp, v)
				if v == u { // v is one of the SCC root. See the video explaination
					break
				}
			}
			scc = append(scc, comp)
		}
	}

	for i, timestamp := range dfn {
		if timestamp == 0 {
			tarjan(i)
		}
	}

	//// https://stackoverflow.com/questions/32750511/does-tarjans-scc-algorithm-give-a-topological-sort-of-the-scc
	slices.Reverse(scc)

	// Tarjan SCC 算法到此结束了。 下面的，把 SCC 缩点变成一个新的 graph
	sid := make([]int, len(g))
	for i, cc := range scc {
		for _, v := range cc {
			sid[v] = i
		}
	}
	ns := len(scc)
	g2 := make([][]int, ns)
	deg := make([]int, ns)
	for v, ws := range g {
		v = sid[v]
		for _, w := range g[v] {
			if v != sid[w] {
				g2[v] = append(g2[v], w)
				deg[w]++
			}
		}
	}
	return scc, sid
}
