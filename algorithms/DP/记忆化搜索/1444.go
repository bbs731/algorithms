package dp

/***


你现在的阶段，是能想出来解题思路的， 但是，证明正确性的能力太差了， 遇到错误的时候，开始动摇然后走到岔路上去了， 这道题，出错的时候，竟然想到， 要再加 两个维度到 dfs(i, j, p, q, k)
来模拟，右下的端点。

dfs(i, j, k)   定义， 左上角 [i,j]  到右下角 [m-1][n-1] 需要切 k 个 pizza 的方案数量。

//枚举 横切，  r = [i+1, m-1]      [i, j] [r-1, n-1] 保证有一个苹果， 且 [r, j] [m-1][n-1] 有  >= k-1 个苹果。    dfs(r, j, k-1)
枚举 竖切,   c = [j+1, n-1]        [i, j] [m-1, c-1] 保证有一个苹果， 且 [i, c] [m-1][n-1] 有  >= k-1 个苹果。    dfs(i, c, k-1)

*/

func ways(pizza []string, k int) int {
	m := len(pizza)
	n := len(pizza[0])
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			if pizza[i][j] == 'A' {
				matrix[i][j] = 1
			}
		}
	}
	sum := NewMatrixSum(matrix)
	cache := make([][][]int, m)
	for i := range cache {
		cache[i] = make([][]int, n)
		for j := range cache[0] {
			cache[i][j] = make([]int, k+1)
			for p := range cache[i][j] {
				cache[i][j][p] = -1
			}
		}
	}

	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if k == 1 {
			if sum.query(i, j, m-1, n-1) > 0 {
				return 1
			}
			return 0
		}
		ans := 0

		if cache[i][j][k] != -1 {
			return cache[i][j][k]
		}

		//枚举 竖切,   c = [j+1, n-1]        [i, j] [m-1, c-1] 保证有一个苹果， 且 [i, c] [m-1][n-1] 有  >= k-1 个苹果。    dfs(i, c, k-1)
		for c := j + 1; c < n; c++ {
			if sum.query(i, j, m-1, c-1) > 0 {
				ans += dfs(i, c, k-1)
			}
		}

		for r := i + 1; r < m; r++ {
			if sum.query(i, j, r-1, n-1) > 0 {
				ans += dfs(r, j, k-1)
			}
		}

		ans = ans % int(1e9+7)
		cache[i][j][k] = ans
		return ans
	}

	return dfs(0, 0, k)
}

type MatrixSum [][]int

func NewMatrixSum(matrix [][]int) MatrixSum {
	m := len(matrix)
	n := len(matrix[0])
	sum := make(MatrixSum, m+1)
	sum[0] = make([]int, n+1)
	// 注意一下顺序， i, j 都是正序
	for i := range matrix {
		sum[i+1] = make([]int, n+1)
		for j := range matrix[0] {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + matrix[i][j]
		}
	}
	return sum
}

/*
**
左上角  [r1,c1] 右下角 [r2,c2]  是闭区间。
*/
func (s MatrixSum) query(r1, c1, r2, c2 int) int {
	return s[r2+1][c2+1] - s[r2+1][c1] - s[r1][c2+1] + s[r1][c1]
}
