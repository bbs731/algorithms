package _600_1900

/*
有 n 个花园，按从 1 到 n 标记。另有数组 paths ，其中 paths[i] = [xi, yi] 描述了花园 xi 到花园 yi 的双向路径。在每个花园中，你打算种下四种花之一。

另外，所有花园 最多 有 3 条路径可以进入或离开.

你需要为每个花园选择一种花，使得通过路径相连的任何两个花园中的花的种类互不相同。

以数组形式返回 任一 可行的方案作为答案 answer，其中 answer[i] 为在第 (i+1) 个花园中种植的花的种类。花的种类用  1、2、3、4 表示。保证存在答案。



示例 1：

输入：n = 3, paths = [[1,2],[2,3],[3,1]]
输出：[1,2,3]
解释：
花园 1 和 2 花的种类不同。
花园 2 和 3 花的种类不同。
花园 3 和 1 花的种类不同。
因此，[1,2,3] 是一个满足题意的答案。其他满足题意的答案有 [1,2,4]、[1,4,2] 和 [3,2,1]
示例 2：

输入：n = 4, paths = [[1,2],[3,4]]
输出：[1,2,1,2]
示例 3：

输入：n = 4, paths = [[1,2],[2,3],[3,4],[4,1],[1,3],[2,4]]
输出：[1,2,3,4]


提示：

1 <= n <= 104
0 <= paths.length <= 2 * 104
paths[i].length == 2

 */

// 感觉就是 simple  flood-fill 不用任何技巧
// 难度 1712 ，你做了1个小时吗？这样肯定不行啊！

func gardenNoAdj(n int, paths [][]int) []int {

	g := make([][]int, n)
	for _, e := range paths {
		u, v := e[0]-1, e[1]-1
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	colors := make([]int, n)
	var dfs func(int)
	dfs = func(v int) {
		//if depth == n {
		//	ans = append([]int{}, colors...)
		//	return true
		//}

		for i := 1; i <= 4; i++ {
			conflict := false
			for _, w := range g[v] {
				if colors[w] == i {
					conflict = true
					break
				}
			}
			if conflict {
				continue
			}
			colors[v] = i
			//if depth == n-1 {
			//	copy(ans, colors)
			//	return true
			//}
			for _, w := range g[v] {
				if colors[w] == 0 {
					dfs(w)
					//if dfs(w, depth+1) {
					//return true
					//}
				}
			}
		}
	}

	for k := range g {
		if colors[k] == 0 {
			dfs(k)
		}
	}
	return colors
}
