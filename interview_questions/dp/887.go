package dp

import "math"

/***
给你 k 枚相同的鸡蛋，并可以使用一栋从第 1 层到第 n 层共有 n 层楼的建筑。

已知存在楼层 f ，满足 0 <= f <= n ，任何从 高于 f 的楼层落下的鸡蛋都会碎，从 f 楼层或比它低的楼层落下的鸡蛋都不会破。

每次操作，你可以取一枚没有碎的鸡蛋并把它从任一楼层 x 扔下（满足 1 <= x <= n）。如果鸡蛋碎了，你就不能再次使用它。如果某枚鸡蛋扔下后没有摔碎，则可以在之后的操作中 重复使用 这枚鸡蛋。

请你计算并返回要确定 f 确切的值 的 最小操作次数 是多少？


示例 1：

输入：k = 1, n = 2
输出：2
解释：
鸡蛋从 1 楼掉落。如果它碎了，肯定能得出 f = 0 。
否则，鸡蛋从 2 楼掉落。如果它碎了，肯定能得出 f = 1 。
如果它没碎，那么肯定能得出 f = 2 。
因此，在最坏的情况下我们需要移动 2 次以确定 f 是多少。
示例 2：

输入：k = 2, n = 6
输出：3
示例 3：

输入：k = 3, n = 14
输出：4

 */

/***
这题， 太难了

f[n][k] = min( max(f[n-x][k] , f[x-1][k-1]))  x from 1 to n
 */

func superEggDrop(k int, n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
		// 初始化
		if i == 1 {
			for eggs := 1; eggs <= k; eggs++ {
				dp[1][eggs] = 1
			}
		}
		dp[i][1] = i
	}

	for i := 2; i <= n; i++ {
		for j := 2; j <= k; j++ {
			dp[i][j] = math.MaxInt32
			// 寻找的是 x ?
			l, r := 1, i
			for l+1 < r {
				mid := (l + r) >> 1
				t1 := dp[i-mid][j]
				t2 := dp[mid-1][j-1]
				if t1 < t2 {
					r = mid
				} else if t1 > t2 {
					l = mid
				} else {
					l = mid
					r = mid
					break
				}
			}
			// l + 1 = r
			// [l, r] 是 candidate x 的值。
			dp[i][j] = min(max(dp[i-l][j], dp[l-1][j-1]), max(dp[i-r][j], dp[r-1][j-1])) + 1
		}
	}
	return dp[n][k]
}
