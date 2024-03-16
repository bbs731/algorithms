package dfs

func countPairs(n int, edges [][]int) int64 {
	g := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	f := make([]bool, n)
	var cnt int
	var dfs func(int)
	dfs = func(i int) {
		cnt++
		f[i] = true

		for _, j := range g[i] {
			if f[j] == false {
				dfs(j)
			}
		}
	}

	l := []int{}
	ans := 0
	sum := 0
	for i := 0; i < n; i++ {
		if f[i] == false {
			cnt = 0
			dfs(i)
			l = append(l, cnt)
			ans = ans + cnt*sum
			sum += cnt
		}
	}
	return int64(ans)
}
