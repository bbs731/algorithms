package dp

func rootCount(edges [][]int, guesses [][]int, k int) int {

	type pair struct {
		x, y int
	}
	g := make([][]int, len(edges)+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	s :=make(map[pair]int, len(guesses))
	for _, q:= range guesses{
		u, v := q[0], q[1]
		s[pair{u, v}] = 1
	}

	cnt0 :=0
	var dfs func(int, int)
	dfs = func(x, fa int){
		for _, y := range g[x] {
			if y != fa {
				if s[pair{x, y}] == 1{
					cnt0++
				}
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)

	ans:=0
	var reroot func(int, int, int)
	reroot = func(x, fa, cnt int) {
		if cnt >= k {
			ans++
		}

		for _, y := range g[x] {
			if y != fa {
				reroot(y, x, cnt - s[pair{x, y}] + s[pair{y, x}])
			}
		}
	}
	reroot(0, -1, cnt0)
	return ans
}
