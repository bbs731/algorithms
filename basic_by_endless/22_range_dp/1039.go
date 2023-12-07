package dp

import "math"

/*
灵神的答案：
https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/solutions/2203005/shi-pin-jiao-ni-yi-bu-bu-si-kao-dong-tai-aty6/

f[i][j] = min(f[i][k] + f[k][j] + v[i][j][k]) for  i < k <j
因为， f[i] 有 f[k] 推倒而来， 然而 k > i 的， 所以， i 的遍历顺序应该是倒序的。
       f[][j] 是由 f[][k] 和 f[][j] 而来， k < j 所以 j 是可以正序枚举的。

这个答案太工整了，
不需要使用额外枚举 length 逐渐变大的枚举技巧。 而且不容易出错。
*/
func minScoreTriangulation(values []int) int {
	n := len(values)

	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			f[i][j] = math.MaxInt
			for k := i + 1; k < j; k++ {
				f[i][j] = min(f[i][j], f[i][k]+f[k][j]+values[i]*values[j]*values[k])
			}
		}
	}
	return f[0][n-1]
}

/*
dfs(i, j) = min (dfs(i, k) + dfs(k, j) + v[i][j][k])  for  i <k <j
f[i][j] = min(f[i][k] + f[k][j] + v[i][j][k]) for  i < k <j
*/
func minScoreTriangulation_chunlei(values []int) int {
	n := len(values)
	inf := math.MaxInt / 2

	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}
	for l := 2; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			f[i][j] = inf
			for k := i + 1; k < j; k++ {
				f[i][j] = min(f[i][j], f[i][k]+f[k][j]+values[i]*values[j]*values[k])
			}
		}
	}
	// 下面这个枚举方式是错的。
	//for i := 0; i < n; i++ {
	//	for j := i + 2; j < n; j++ {
	//		ans := inf
	//		for k := i + 1; k < j; k++ {
	//			ans = min(ans, f[i][k]+f[k][j]+values[i]*values[j]*values[k])
	//		}
	//		f[i][j] = ans
	//	}
	//}
	return f[0][n-1]
}

/*
读懂题的意思还是挺有难度的。

dfs(i, j) = min (dfs(i, k) + dfs(k, j) + v[i][j][k])  for  i <k <j

 */
func minScoreTriangulation_dfs(values []int) int {
	inf := int(1e10)
	n := len(values)
	var dfs func(int, int) int
	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	dfs = func(i, j int) int {
		if i+2 > j {
			return 0
		}

		if cache[i][j] != -1 {
			return cache[i][j]
		}
		ans := inf
		for k := i + 1; k < j; k++ {
			ans = min(ans, dfs(i, k)+dfs(k, j)+values[i]*values[j]*values[k])
		}
		cache[i][j] = ans
		return ans
	}
	return dfs(0, n-1)
}
