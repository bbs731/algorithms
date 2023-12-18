package _3_tree_dp_diameter

/*
给你一棵 树（即一个连通、无向、无环图），根节点是节点 0 ，这棵树由编号从 0 到 n - 1 的 n 个节点组成。用下标从 0 开始、长度为 n 的数组 parent 来表示这棵树，其中 parent[i] 是节点 i 的父节点，由于节点 0 是根节点，所以 parent[0] == -1 。

另给你一个字符串 s ，长度也是 n ，其中 s[i] 表示分配给节点 i 的字符。

请你找出路径上任意一对相邻节点都没有分配到相同字符的 最长路径 ，并返回该路径的长度。

 */
/*
https://leetcode.cn/problems/longest-path-with-different-adjacent-characters/solutions/1427611/by-endlesscheng-92fw/
灵神给的答案， 算的是边的个数。
*/

func longestPath(parent []int, s string) int {
	ans := 1
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		g[parent[i]] = append(g[parent[i]], i)
	}
	var dfs func(root int) int
	dfs = func(root int) int {
		var p int
		//for i := 1; i < n; i++ {  //  优化一下这里. 通过 预处理 parent and children list
		for _, i := range g[root] {
			//if parent[i] == root {
			cl := dfs(i) // get children path's length  这里算的是 Node 的个数
			if s[i] != s[root] {
				ans = max(ans, cl+p+1)
				p = max(p, cl)
			}
			//}
		}
		return p + 1
	}
	dfs(0)
	return ans
}
