package chagne_root_DP

/*
给定一个无向、连通的树。树中有 n 个标记为 0...n-1 的节点以及 n-1 条边 。
给定整数 n 和数组 edges ， edges[i] = [ai, bi]表示树中的节点 ai 和 bi 之间有一条边。
返回长度为 n 的数组 answer ，其中 answer[i] 是树中第 i 个节点与所有其他节点之间的距离之和。
*/
/*
看看灵神的答案启发良多：
https://leetcode.cn/problems/sum-of-distances-in-tree/solutions/2345592/tu-jie-yi-zhang-tu-miao-dong-huan-gen-dp-6bgb/

	得到的递推式子：
	ans[y] = ans[x] + n - size[y] - size[y]
	y 是 x 的 child, size[i] 代表 node i 的子树的大小(包括自己）
 */
func sumOfDistancesInTree(n int, edges [][]int) []int {

	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	ans := make([]int, n)
	size := make([]int, n)
	//cost := 0  直接用 ans[0] 就行
	var dfs func(int, int, int) int
	dfs = func(c int, fa int, depth int) int {
		cnt := 1 //算 size[c] 的时候, 把自己需要算进去，所以初始化为 1
		for _, e := range g[c] {
			if e != fa {
				cnt += dfs(e, c, depth+1)
			}
		}
		size[c] = cnt
		//cost += depth
		ans[0] += depth
		return size[c]
	}
	// 第一遍 dfs 得到每个子树的 size
	dfs(0, -1, 0)
	//ans[0] = cost

	var reroot func(int, int)
	reroot = func(x int, fa int) {
		// 把这段写在下面，更加的简洁, 这样就不需要判断 fa != -1 了。
		//if fa != -1 {
		//	ans[c] = ans[fa] + n - size[c]*2
		//}
		for _, y := range g[x] {
			if y != fa {
				ans[y] = ans[x] + n - size[y]*2
				reroot(y, x)
			}
		}
	}
	reroot(0, -1)
	return ans
}
