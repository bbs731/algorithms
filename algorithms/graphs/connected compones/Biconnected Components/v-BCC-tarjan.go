package one_day_exercise

//时间复杂度线性的 O(n+m） 能同时求出:  无向图中双连通分量 + 割点
//缩点之后，变成一颗树对吗？ 缩点的代码，在哪儿有？
// 这个代码，是从 李煜东 <<算法竞赛进阶指南>> 里扒出来的。
// 感觉，比灵神， v-bcc 的代码简洁一点(tarjan 里不带 fa 参数）。 但对于 root 的用法不是很确定, 但是我感觉是对的, 把root 用作变量，不是 const。


func bccTarjan( g [][]int) (dcc [][]int, isCut []bool){
	var root int
	dcc = [][]int{}
	isCut = make([]bool, len(g))

	dfn := make([]int, len(g))
	low := make([]int, len(g))
	dfsClock := 0
	cnt := 0
	st := []int{}


	var tarjan func(int)
	tarjan = func(u int) {
		dfsClock++
		dfn[u]= dfsClock
		low[u] = dfsClock
		if u == root && len(g[u]) == 0 {  // 对于 root 就是孤立的一个点。
			dcc = append(dcc, []int{u})
			return
		}
		st = append(st, u)
		childCnt := 0

		for _, v := range g[u]{
			if dfn[v] == 0 {
				tarjan(v)
				low[u] = min(low[u], low[v])

				if low[v] >= dfn[u]{
					childCnt++
					if u != root || childCnt> 1 { // 如果是 root childCnt >= 2 才是 cut
						isCut[u]= true
					}
					cnt++
					comp := []int{}
					for {
						z := st[len(st)-1]
						st = st[:len(st)-1]
						comp = append(comp, z)
						if z == v {
							break
						}
					}
					comp = append(comp, u)
					dcc = append(dcc, comp)
				}
			} else {
				low[u]= min(low[u], dfn[v])
			}
		}
	}

	//tarjan(root)

	for v, timestamp := range dfn {
		if timestamp == 0 {
			root = v
			tarjan(v)
		}
	}
	return
}
