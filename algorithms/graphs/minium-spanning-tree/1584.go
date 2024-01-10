package minium_spanning_tree

func calDist(p1, p2 []int) int {
	abs := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	return abs(p1[0], p2[0]) + abs(p1[1], p2[1])
}

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	dis := make([][]int, n)
	for i := 0; i < n; i++ {
		dis[i] = make([]int, n)
	}
	const inf int = 1e9

	for i, p1 := range points {
		dis[i][i] = inf
		for j := i + 1; j < n; j++ {
			p2 := points[j]
			dis[i][j] = calDist(p1, p2)
			dis[j][i] = dis[i][j]
		}
	}

	sum, _ := mstPrim(dis, 0)
	return sum
}

func mstPrim(dis [][]int, root int) (mstSum int, edges [][2]int) {
	edges = make([][2]int, 0, len(dis)-1)
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
