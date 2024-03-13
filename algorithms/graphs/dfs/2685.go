package dfs

func countCompleteComponents(n int, edges [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	var v, e int

	f := make([]bool, n)
	var dfs func(int)
	dfs = func(u int) {
		f[u] = true
		v++
		for _, v := range g[u] {
			e += 1
			if f[v] == false {
				dfs(v)
			}
		}
	}

	var cnt int
	for i := 0; i < n; i++ {
		if f[i] == false {
			v, e = 0, 0
			dfs(i)
			//fmt.Println(v, e)
			if v*(v-1) == e {
				cnt++
			}
		}
	}
	return cnt
}
