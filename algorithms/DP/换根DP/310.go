package chagne_root_DP

//A tree is an undirected graph in which any two vertices are connected by exact
//ly one path. In other words, any connected graph without simple cycles is a tree
//.
//
// Given a tree of n nodes labelled from 0 to n - 1, and an array of n - 1 edges
// where edges[i] = [ai, bi] indicates that there is an undirected edge between th
//e two nodes ai and bi in the tree, you can choose any node of the tree as the ro
//ot. When you select a node x as the root, the result tree has height h. Among al
//l possible rooted trees, those with minimum height (i.e. min(h)) are called mini
//mum height trees (MHTs).
//
// Return a list of all MHTs' root labels. You can return the answer in any orde
//r.
//
// The height of a rooted tree is the number of edges on the longest downward pa
//th between the root and a leaf.
//
//
// Example 1:
//
//
//Input: n = 4, edges = [[1,0],[1,2],[1,3]]
//Output: [1]
//Explanation: As shown, the height of the tree is 1 when the root is the node w
//ith label 1 which is the only MHT.
//
//
// Example 2:
//
//
//Input: n = 6, edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]
//Output: [3,4]
//
//
// Constraints:
//
//
// 1 <= n <= 2 * 104
// edges.length == n - 1
// 0 <= ai, bi < n
// ai != bi
// All the pairs (ai, bi) are distinct.

/*
换根 DP， 这道题，真好玩！ 一个 medium 竟然是所有 hard re-root DP  里面，写的最复杂的那个！
 */
func findMinHeightTrees(n int, edges [][]int) []int {
	g := make([][]int, n)
	// process edges to graphs
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	heights := make(map[int][2]int) // heights[i] 存 node i 的两个最高子树的长度。
	var dfs func(int, int) int      // 返回树的高度。
	dfs = func(x, fa int) int {
		for _, y := range g[x] {
			if y != fa {
				h := dfs(y, x) + 1
				if h >= heights[x][1] {
					//heights[x][0], heights[x][1] = heights[x][1], h  // 这里不能直接赋值，golang 的编译器不允许！ 哈哈，头一次见到。
					heights[x] = [2]int{heights[x][1], h}
				} else {
					//heights[x][0] = max(heights[x][0], h)
					heights[x] = [2]int{max(heights[x][0], h), heights[x][1]}
				}
			}
		}
		return heights[x][1]
	}
	dfs(0, -1)
	// now we have 0 root tree, and each subtree has its top 2 heights

	mh := n
	ans := make([]int, 0)

	/*
	 怀疑是不是想的太难了， 不知道有没有更简单的方法， 搜一搜灵神的答案：

	这第二遍 re-root 的过程， 需要两个边界来收敛， 只用 long 是不对的， 因为 short 那边会增长直到超过 long 这边的树的高度。
		1. 只需要 按照最长的子树去搜索就可以了， 因为，其它的路径，只会让树高变得更长。
		2. short 是递增数组，所以可以用来剪枝, 譬如 当  max(short, long) > mh 的时候，就没必要再搜索了，可以提前返回了。
	 */
	var reroot func(int, int, int, int)
	reroot = func(x, fa, short, long int) {
		if max(short, long) < mh {
			mh = max(short, long)
			ans = []int{x}
			//return
		} else if max(short, long) == mh {
			ans = append(ans, x)
		} else {
			// 这里，属于剪枝。 再往下搜索不可能再有好结果了。 因为 short 是递增的数组。
			return
		}

		for _, y := range g[x] {
			if y != fa {
				if heights[x][1] == heights[y][1]+1 { // 只有最长的那个 子树 去 dfs 才有意义。 想想为什么， 其它的子树， height 只会更大，没必要再去搜索了，这相当于剪枝了。
					reroot(y, x, max(short+1, heights[x][0]+1), max(heights[y][1], heights[x][0]+1))
				}
			}
		}
	}
	reroot(0, -1, heights[0][0], heights[0][1])
	return ans
}
