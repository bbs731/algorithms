package dfs

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	g := make([][]int, n)

	/****
		这题还容易错了， 爆炸的规则太复杂了， 不是对称的（哎，想当然了！）
	 */
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			p, q := bombs[i], bombs[j]
			if i != j && (q[0]-p[0])*(q[0]-p[0])+(q[1]-p[1])*(q[1]-p[1]) <= p[2]*p[2] {
				g[i] = append(g[i], j)
				//g[j] = append(g[j], i)
			}
		}
	}
	ans := 0
	for i := range g {
		vis := make([]bool, n)
		var dfs func(int)
		cnts := 0
		dfs = func(x int) {
			vis[x] = true
			cnts++
			for _, v := range g[x] {
				if vis[v] == false {
					dfs(v)
				}
			}
		}

		if vis[i] == false {
			cnts = 0
			dfs(i)
			ans = max(ans, cnts)
		}
	}
	return ans
}
