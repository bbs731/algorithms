package dp

import "math"

/*

很好的题， 不是道过一段时间之后，能不能作对。和1039 使用的技巧类似。

	dfs(i, j) = min ( max(dfs(i, k-1), dfs(k+1,j) +  k )  for i <k <j
边界条件：
	dfs(i, j) = 0 if i >=j
	dfs(i, i+1) = i


	f[i][j] = min( max( f[i][k-1],  f[k+1][j] )  + k )  for  i < k < j
初始化：
	f[i][j] = 0  if i >=j
	f[i][i+1] = i

	因为 f[i][] 依赖  f[k][]  k > i 所以  i 的遍历应该是逆序的。 j 的遍历是正序的。
 */

func getMoneyAmount(n int) int {

	f := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]int, n+1)
		if i < n { // 这里 的边界条件注意， 要不然 i+1会越界
			f[i][i+1] = i
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n; j++ { // 这里的 j 的初始值需要注意， 因为如果 j = i+1  那么 f[i][i+1] = inf 会覆盖 初始化的值 i
			f[i][j] = math.MaxInt
			for k := i + 1; k < j; k++ {
				f[i][j] = min(f[i][j], max(f[i][k-1], f[k+1][j])+k)
			}
		}
	}
	return f[1][n]
}

/*
	dfs(i, j) = min ( max(dfs(i, k-1), dfs(k+1,j) +  k )  for i <k <j
	// dfs(i, k-1), dfs(k+1, j)   k-1 这个状态zhangchunlei 第一次就错了。
边界条件：
	dfs(i, j) = 0 if i >=j
	dfs(i, i+1) = i
 */
func getMoneyAmount_dfs(n int) int {

	cache := make([][]int, n+1) // 开 n+1
	for i := range cache {
		cache[i] = make([]int, n+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i >= j {
			return 0
		}
		if i+1 == j {
			return i
		}

		if cache[i][j] != -1 {
			return cache[i][j]
		}

		ans := math.MaxInt
		for k := i + 1; k < j; k++ {
			ans = min(ans, max(dfs(i, k-1), dfs(k+1, j))+k)
		}
		cache[i][j] = ans
		return ans
	}
	return dfs(1, n)
}
