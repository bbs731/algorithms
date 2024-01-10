package Euler_Tour

import (
	"math"
	"strconv"
)

/*
这道题， 最难的是如何，想象这个图是什么样子的。

k^(n-1) 个节点。 组成的图， 恰好，全连接，每个节点，k in edges, k out edges 组成的图。

这张图的，euler 遍历，所有的边，拼在一起就是答案。
 */
func crackSafe(n int, k int) string {

	mod := int(math.Pow(10, float64(n-1)))
	visited := make(map[int]bool)

	ans := ""
	var dfs func(int)
	dfs = func(v int) {
		for i := 0; i < k; i++ {
			w := v*10 + i
			if !visited[w] {
				visited[w] = true
				dfs(w % mod)
				// 加边
				ans = ans + strconv.Itoa(i)
			}
		}
	}

	dfs(0)
	for i := 1; i < n; i++ {
		ans += "0"
	}
	return ans
}
