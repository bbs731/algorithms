package dfs

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	out := make([]int, n)
	gf := make([][]int, n)

	for i, g := range graph {
		for _, p := range g {
			out[i]++
			gf[p] = append(gf[p], i)
		}
	}
	ans := make([]int, 0)

	q := []int{}
	for i := 0; i < n; i++ {
		if out[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		tmp := q
		q = nil
		ans = append(ans, tmp...)

		for _, t := range tmp {
			for _, p := range gf[t] {
				out[p]--
				if out[p] == 0 {
					q = append(q, p)
				}
			}
		}
	}
	sort.Ints(ans)
	return ans
}
