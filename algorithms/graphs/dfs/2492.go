package dfs

import "math"

func minScore(n int, roads [][]int) int {
	type pair struct {
		to, d int
	}
	g := make([][]pair, n+1)
	for _, r := range roads {
		a, b, d := r[0], r[1], r[2]
		g[a] = append(g[a], pair{b, d})
		g[b] = append(g[b], pair{a, d})
	}

	f := make([]bool, n+1)
	var dfs func(int)
	dfs = func(i int) {
		f[i] = true
		for _, p := range g[i] {
			to := p.to
			if f[to] == false {
				dfs(to)
			}
		}
	}
	dfs(1)

	ans := math.MaxInt32

	for _, r := range roads {
		a, b, d := r[0], r[1], r[2]
		if f[a] == true && f[b] == true {
			ans = min(ans, d)
		}
	}
	return ans
}

func minScore(n int, roads [][]int) int {
	type pair struct {
		to, d int
	}
	g := make([][]pair, n+1)
	for _, r := range roads {
		a, b, d := r[0], r[1], r[2]
		g[a] = append(g[a], pair{b, d})
		g[b] = append(g[b], pair{a, d})
	}

	ans := math.MaxInt32
	f := make([]bool, n+1)
	var dfs func(int)
	dfs = func(i int) {
		f[i] = true
		for _, p := range g[i] {
			to := p.to
			ans = min(ans, p.d)
			if f[to] == false {
				dfs(to)
			}
		}
	}
	dfs(1)

	return ans
}
