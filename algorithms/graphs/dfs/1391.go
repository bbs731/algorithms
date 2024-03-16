package dfs

func makeConnected(n int, connections [][]int) int {
	g := make([][]int, n)
	m := len(connections)
	for _, c := range connections {
		a, b := c[0], c[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	conn := 0
	f := make([]bool, n)
	// 用 dfs 发现连通块
	var dfs func(int)
	dfs = func(x int) {
		f[x] = true
		for _, j := range g[x] {
			if f[j] == false {
				dfs(j)
			}
		}
	}

	for i := 0; i < n; i++ {
		if f[i] == false {
			conn++
			dfs(i)
		}
	}
	// 考察的就是这段逻辑吧， 你思考了好几遍啊！
	if m >= n-1 {
		return conn - 1
	}
	return -1
}
