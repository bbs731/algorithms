package minium_spanning_tree

import "sort"

//https://oi-wiki.org/graph/mst/

// implementation of Prim MST 当输入是稠密图的时候。
//代码来自灵神：
//https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/graph.go#L1969
type graph struct{}

// 因为是朴素的找最小值，所以时间复杂度是  O(n^2 + m)
func (*graph) mstPrim(dis [][]int, root int) (mstSum int, edges [][2]int) {
	edges = make([][2]int, 0, len(dis)-1) // 因为是树，所以， num of edges = n-1
	// 注意：dis 需要保证 dis[i][i] = inf，从而避免自环的影响

	const inf int = 1e9
	// minD[i].d 表示当前 MST 到点 i 的最小距离，对应的边为 minD[i].v-i
	minD := make([]struct{ v, d int }, len(dis))
	for i := range minD {
		minD[i].d = inf
	}
	minD[root].d = 0
	inMST := make([]bool, len(dis))

	for {
		// 根据切分定理，求不在当前 MST 的点到当前 MST 的最小距离，即 minD[v].d
		v := -1
		for w, in := range inMST {
			if !in && (v < 0 || minD[w].d < minD[v].d) {
				v = w
			}
		}
		if v < 0 {
			// 已求出 MST
			return
		}

		// 加入 MST
		inMST[v] = true
		mstSum += minD[v].d

		if v != root {
			edges = append(edges, [2]int{minD[v].v, v})
		}

		// 更新 minD
		for w, d := range dis[v] { // 因为稠密图才用 Prim, 所以 dis 是 adjacency matrix
			if !inMST[w] && d < minD[w].d {
				minD[w].d = d
				minD[w].v = v
			}
		}
	}
}

func (*graph) Kruskal(n int, edges [][]int) int {
	sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })

	fa := make([]int, n) // 并查集
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

	sum := 0
	cntE := 0
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]
		fv, fw := find(v), find(w)
		if fv != fw {
			fa[fv] = fw
			sum += wt
			cntE++
		}
	}

	if cntE < n-1 {
		return -1
	}

	return sum
}
