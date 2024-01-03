package topological_sort

// Kahn 求拓扑排序的例子。 Tree 是一个天然的试验场。
//https://oi-wiki.org/graph/topo/



//题解： https://leetcode.cn/problems/collect-coins-in-a-tree/solutions/2191371/tuo-bu-pai-xu-ji-lu-ru-dui-shi-jian-pyth-6uli/
//虽然 Topological Sort 只适用于有向图 DAG。 而且Tree 不是有向图，但是Tree是无环的, 也可以使用 Topological Sort
//可以一层一层的摘掉叶子节点，就是 Topological Sort Order (Kahn 算法 https://oi-wiki.org/graph/topo/）
//Video： https://www.youtube.com/watch?v=cIBFEhD77b4&list=PLDV1Zeh2NRsDGO4--qE8yH72HFL1Km93P&index=16

func collectTheCoins(coins []int, edges [][]int) int {
	n := len(coins)
	deg := make([]int, n) // degree
	g := make([][]int, n)
	leftEdges := n - 1

	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		deg[u]++
		g[v] = append(g[v], u)
		deg[v]++
	}

	q := []int{}
	for i := 0; i < n; i++ {
		if deg[i] == 1 && coins[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		t := q[len(q)-1]
		deg[t]--
		leftEdges--
		q = q[:len(q)-1]

		for _, e := range g[t] {
			deg[e]--
			if deg[e] == 1 && coins[e] == 0 {
				q = append(q, e)
			}
		}
	}

	// 这次是找虽有的有金币的叶子节点。
	q = []int{}
	for i := 0; i < n; i++ {
		if deg[i] == 1 && coins[i] > 0 {
			q = append(q, i)
		}
	}

	leftEdges -= len(q)

	for len(q) > 0 {
		t := q[len(q)-1]
		deg[t]--
		q = q[:len(q)-1]

		for _, e := range g[t] {
			deg[e]--

			//去掉所有叶子，然后再去掉新产生的叶子，剩余节点就是必须要访问的节点。 不需要考虑金币的问题。
			//if deg[e] == 1 && coins[e] == 0 {
			if deg[e] == 1 { // 这里为什么不用再比较 coins[e] == 0 了呢？ // 只要是新产生的叶子就可以。
				leftEdges--
			}
		}
	}

	return max(leftEdges*2, 0)
}
