package minium_spanning_tree

import (
	"sort"
)

/*

输入：n = 5, edges = [[0,1,1],[1,2,1],[2,3,2],[0,3,2],[0,4,3],[3,4,3],[1,4,6]]
输出：[[0,1],[2,3,4,5]]
解释：上图描述了给定图。
下图是所有的最小生成树。
 */

// Kruskal + Tarjan 结合的算法， 全懂了， 2024.1.11  不知道以后会不会又犯糊涂
// union-find 最简单的路径压缩就可以，不用官方的考虑 size 就可以， 但是 union 的时候应该 union fv, fw 这里容易出错。
// 清理 tarjan 的图的时候容易犯错， 应该清理 fv, fw 而不是 v,w  或者直接 clear g, dfn 更方便不容易出错。

// https://oi-wiki.org/graph/mst/   中判断 MST 是否唯一的例子，借不了这道题。
// 换一个枚举的笨方法。

// 另外伪边就是应为， weight 相等的边造成的。 Kruskal 算法有一个性质：
// 在 处理完 所有 <= W weight 的 edges 之后， union-find set 的连通性是唯一的。
// 那么，可以，判断所有 = W weight 的 edges, 那些 edge 是割边，他就是CritialEdge 否则 就是 PseudoCritical Edge （当然在排除，不可能加入 MST 的 edge 之后）

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	must := make([]int, 0)
	non_must := make([]int, 0)

	marks := mstKruskal(n, edges)
	for i, m := range marks {
		if m == 1 {
			must = append(must, i)
		} else if m == 0 {
			non_must = append(non_must, i)
		}
	}

	return [][]int{must, non_must}
}

func mstKruskal(n int, edges [][]int) (marks []int) {
	for i, e := range edges {
		edges[i] = append(e, i) // 把边的 index 放在 edges[i][3] 的位置。
	}
	// 边权范围小的话也可以用桶排
	sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })

	fa := make([]int, n) // n+1
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	type neighbour struct{ to, index int }
	g := make([][]neighbour, n)
	m := len(edges)
	marks = make([]int, m)
	dfn := make([]int, n)
	dfsClock := 0
	low := make([]int, n)

	var tarjan func(int, int)
	tarjan = func(u, fid int) {
		dfsClock++
		low[u] = dfsClock
		dfn[u] = dfsClock

		for _, e := range g[u] {
			if v := e.to; dfn[v] == 0 {
				tarjan(v, e.index)
				low[u] = min(low[u], low[v])
				if low[v] > dfn[u] {
					marks[e.index] = 1 // bridge cut
				}

			} else if e.index != fid {
				low[u] = min(low[u], dfn[v])
			}
		}
	}

	//uf := newUnionFind(n)
	for i := 0; i < m; {
		j := i
		vs := []int{}
		for ; j < m && edges[j][2] == edges[i][2]; j++ {
			v, w := find(edges[j][0]), find(edges[j][1])
			//v, w := uf.find(edges[j][0]), uf.find(edges[j][1])
			if v == w {
				marks[edges[j][3]] = -1 // edge not in MST
			} else {
				g[v] = append(g[v], neighbour{w, edges[j][3]})
				g[w] = append(g[w], neighbour{v, edges[j][3]})
				vs = append(vs, v, w)
			}
		}

		// run tarjan to find bridge cut
		for _, u := range vs {
			if dfn[u] == 0 {
				tarjan(u, -1)
			}
		}

		// real Kruskal to add edge to MST
		for k := i; k < j; k++ {
			e := edges[k]
			v, w := e[0], e[1]
			fv, fw := find(v), find(w)
			if fv != fw {
				fa[fv] = fw
			}
			//g[v] = nil  // 这里容易犯错误， 因为之前建图是用 u, v 在 union-find 中的 root建图的，所以清理的时候要小心， g[v] 是不对的。
			// 清理的时候也 应该清理 g[fv], g[fw]
			g[fv] = nil
			g[fw] = nil
			dfn[fv] = 0
			dfn[fw] = 0
		}
		// reset tarjan used states
		// 当然， 直接 clear g, dfn, low 更方便一些，而且不出错。
		//clear(g)
		//clear(dfn)
		//clear(low)
		//dfsClock = 0

		// next round
		i = j
	}
	return
}
