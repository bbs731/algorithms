package weekly

import "math"

// 灵神的答案。 太牛了。
// 需要总结的点。
// 通过 dfsClock 可以知道， x, y 的祖先关系。

// https://leetcode.cn/problems/minimum-score-after-removals-on-a-tree/solutions/1625899/dfs-shi-jian-chuo-chu-li-shu-shang-wen-t-x1kk/

func minimumScore(nums []int, edges [][]int) int {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	xor := make([]int, n)
	in := make([]int, n)
	out := make([]int, n)
	clock := 0
	var dfs func(int, int)
	dfs = func(x, fa int) {
		clock++
		in[x] = clock
		xor[x] = nums[x]
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x)
				xor[x] ^= xor[y]
			}
		}
		out[x] = clock
	}
	dfs(0, -1)

	//这个，第一次学习到。
	isAncestor := func(x, y int) bool { return in[x] < in[y] && in[y] <= out[x] }

	ans := math.MaxInt32
	// 枚举点， 不是 edges, 这个也是没想到。
	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			var x, y, z int
			if isAncestor(i, j) { // i 是 j 的祖先节点
				x, y, z = xor[j], xor[i]^xor[j], xor[0]^xor[i]
			} else if isAncestor(j, i) { // j 是 i 的祖先节点
				x, y, z = xor[i], xor[i]^xor[j], xor[0]^xor[j]
			} else { // 删除的两条边分别属于两颗不相交的子树
				x, y, z = xor[i], xor[j], xor[0]^xor[i]^xor[j]
			}
			ans = min(ans, max(max(x, y), z)-min(min(x, y), z))
			if ans == 0 {
				return 0 // 提前退出
			}
		}
	}
	return ans
}
