package chagne_root_DP

/*

Alice 有一棵 n 个节点的树，节点编号为 0 到 n - 1 。树用一个长度为 n - 1 的二维整数数组 edges 表示，其中 edges[i] = [ai, bi] ，表示树中节点 ai 和 bi 之间有一条边。

Alice 想要 Bob 找到这棵树的根。她允许 Bob 对这棵树进行若干次 猜测 。每一次猜测，Bob 做如下事情：

选择两个 不相等 的整数 u 和 v ，且树中必须存在边 [u, v] 。
Bob 猜测树中 u 是 v 的 父节点 。
Bob 的猜测用二维整数数组 guesses 表示，其中 guesses[j] = [uj, vj] 表示 Bob 猜 uj 是 vj 的父节点。

Alice 非常懒，她不想逐个回答 Bob 的猜测，只告诉 Bob 这些猜测里面 至少 有 k 个猜测的结果为 true 。

给你二维整数数组 edges ，Bob 的所有猜测和整数 k ，请你返回可能成为树根的 节点数目 。如果没有这样的树，则返回 0。


示例 1：



输入：edges = [[0,1],[1,2],[1,3],[4,2]], guesses = [[1,3],[0,1],[1,0],[2,4]], k = 3
输出：3
解释：
根为节点 0 ，正确的猜测为 [1,3], [0,1], [2,4]
根为节点 1 ，正确的猜测为 [1,3], [1,0], [2,4]
根为节点 2 ，正确的猜测为 [1,3], [1,0], [2,4]
根为节点 3 ，正确的猜测为 [1,0], [2,4]
根为节点 4 ，正确的猜测为 [1,3], [1,0]
节点 0 ，1 或 2 为根时，可以得到 3 个正确的猜测。


[[0,1],[2,0],[0,3],[4,2],[3,5],[6,0],[1,7],[2,8],[2,9],[4,10],[9,11],[3,12],[13,8],[14,9],[15,9],[10,16]]

[[8,2],[12,3],[0,1],[16,10]]
 */

/*

这是一道好题， 通过 换根 DP 找到了编程的乐趣。又因为 bug 失去了，刚找到的乐趣 :)

思路：
1. 典型的 换根 DP 的问题， 那么需要两次 dfs 来求解，那么，这个就是这个程序的主干。
2. 我们需要统计变化的数据， 很明显，就是 guesses 的对错结果。 用和个 hash table 把 guesses 的结果保存起来， 然后换 root 的时候，我们就根据更新的变化，重新统计 guesses 的结果。把满足要求的 ans 记录下来。
 */
func rootCount(edges [][]int, guesses [][]int, k int) int {
	n := len(edges) + 1 // trees have n-1 edges

	// 处理下 guesses, 用 hashtable 存储，方便查询。
	hg := make(map[[2]int]struct{})
	for _, g := range guesses {
		hg[[2]int{g[0], g[1]}] = struct{}{} // 只需要查询， (x,y) 在或者不在 queries 里面。
	}
	//total_correct := 0
	ans := 0
	correct := make([]int, n)

	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int)
	dfs = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				if _, ok := hg[[2]int{x, y}]; ok {
					//total_correct++
					correct[0]++
				}
				dfs(y, x)
			}
		}
	}
	dfs(0, -1) // node 0 as root
	//correct[0] = total_correct

	var reroot func(int, int)
	reroot = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				// use y as new root
				correct[y] = correct[x]
				if _, ok := hg[[2]int{x, y}]; ok {
					correct[y] -= 1
				}
				if _, ok := hg[[2]int{y, x}]; ok {
					//total_correct++
					correct[y] += 1
				}
				// 直接用 total_correct 是不对的，如果需要这样使用 total correct 不应该作为global variable, 应该作为参数，传入 reroot func 中。
				//if total_correct >= k {
				//	ans++
				//}
				reroot(y, x)
			}
		}
	}
	reroot(0, -1)

	for _, v := range correct {
		if v >= k {
			ans++
		}
	}
	return ans

}

func rootCount(edges [][]int, guesses [][]int, k int) int {
	n := len(edges) + 1 // trees have n-1 edges

	// 处理下 guesses, 用 hashtable 存储，方便查询。
	hg := make(map[[2]int]struct{})
	for _, g := range guesses {
		hg[[2]int{g[0], g[1]}] = struct{}{} // 只需要查询， (x,y) 在或者不在 queries 里面。
	}
	total_correct := 0
	ans := 0

	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int)
	dfs = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				if _, ok := hg[[2]int{x, y}]; ok {
					total_correct++
				}
				dfs(y, x)
			}
		}
	}
	dfs(0, -1) // node 0 as root
	//correct[0] = total_correct
	//if total_correct >= k {
	//	ans++
	//}

	var reroot func(int, int, int)
	reroot = func(x, fa, cnt int) {
		if cnt >= k {
			ans++
		}
		for _, y := range g[x] {
			// 这是循环，需要重复利用 cnt, 所以需要copy 一份。
			copy := cnt // 这里不 copy 直接使用 cnt 会掉坑里面，哎！
			if y != fa {
				// use y as new root
				if _, ok := hg[[2]int{x, y}]; ok {
					copy--
				}
				if _, ok := hg[[2]int{y, x}]; ok {
					//total_correct++
					copy++
				}
				// 直接用 total_correct 是不对的，如果需要这样使用 total correct 不应该作为global variable, 应该作为参数，传入 reroot func 中。
				//if total_correct >= k {
				reroot(y, x, copy)
			}
		}
	}
	reroot(0, -1, total_correct)
	return ans
}
