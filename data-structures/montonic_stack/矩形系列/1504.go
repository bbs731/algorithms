package monotonic_stack

/***

给你一个 m x n 的二进制矩阵 mat ，请你返回有多少个 子矩形 的元素全部都是 1 。


示例 1：


输入：mat = [[1,0,1],[1,1,0],[1,1,0]]
输出：13
解释：
有 6 个 1x1 的矩形。
有 2 个 1x2 的矩形。
有 3 个 2x1 的矩形。
有 1 个 2x2 的矩形。
有 1 个 3x1 的矩形。
矩形数目总共 = 6 + 2 + 3 + 1 + 1 = 13 。
示例 2：



输入：mat = [[0,1,1,0],[0,1,1,1],[1,1,1,0]]
输出：24
解释：
有 8 个 1x1 的子矩形。
有 5 个 1x2 的子矩形。
有 2 个 1x3 的子矩形。
有 4 个 2x1 的子矩形。
有 2 个 2x2 的子矩形。
有 2 个 3x1 的子矩形。
有 1 个 3x2 的子矩形。
矩形数目总共 = 8 + 5 + 2 + 4 + 2 + 2 + 1 = 24 。


1 <= m, n <= 150
mat[i][j] 仅包含 0 或 1

 */

/***
基于 84 的应用。 这个矩形的统计，做的非常的秒啊！
 */
 func s84 (heights []int) int {
	n :=len(heights)
	left := make([]int, n)
	right := make([]int, n)
	for i := range right {
		right[i]= n
	}

	st := []int{-1}
	for i, v := range heights {
		for len(st)	 > 1 && v < heights[st[len(st)-1]] {
			right[st[len(st)-1]] = i
			// pop stack
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}

	ans := 0
	for i, v := range heights {
		//ans = max(ans, (right[i] - left[i]-1)*v)
		//计算结果时, 乘法原理 * 矩阵高度    // 这是为什么呢？
		ans += (right[i]-i) *(i-left[i])*v
	}
	return ans
}


func numSubmat(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])

	heights := make([]int, n)
	ans := 0

	for i:=0; i<m; i++ { // 一行一行的来
		for j:=0; j<n; j++ {
			if mat[i][j] == 0 {
				heights[j] = 0  // 这个不连续，直接清零
			} else {
				heights[j]++   // 这个累加
			}
		}
		// now we have heights, 我们用 84题的解法
		ans += s84(heights)
	}
	return ans
}
