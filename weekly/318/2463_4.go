package weekly

import (
	"math"
	"sort"
)

/***
	// dfs[i][j]  代表， [i, n-1] 的工厂，修理 [j, m-1] 个机器人的最小值。
	// i 是工厂，  j 是机器人的下标。
	dfs[i][j] = dfs[i+1][j] ,  dfs[i+1][j+k] + sumDist(j, j+k-1)   for all possible k


	// 需要和 dfs 反着定义
	f[i][j] 代表前 i 个工厂， 修理了 j 的机器人的最小值。
	f[i][j] = min (f[i-1][j],  f[i-1][j-k] + sumDist[1, k]  for all possible k
	// 这里需要正着枚举 i, 倒着枚举 j,  为什么？ 因为， 如果正着枚举j 的话， f[i-1][j-k] 的值会被 f[i][j-k] 的值覆盖掉了。

	f[i+1][j+1] = min( f[i][j+1], f[i][j-k+1] + sumDist[1, k] for all possible k

初始化：
	f[0][0] = 1
return
	f[n][m]
 */

func minimumTotalDistance(robot []int, f[][]int) int64 {
	sort.Ints(robot)
	sort.Slice(f, func(i, j int) bool {
		return f[i][0] < f[j][0]
	} )

	n, m := len(f), len(robot)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = 1e18
		}
	}
	// 这个初始化重要。
	for i:=0; i<=n; i++ {
		dp[i][0] = 0
	}
	for i:=0; i<n; i++ { // factory
		for j:=m-1; j>=0; j-- {
			dp[i+1][j+1] = dp[i][j+1]  // ith厂子不修理 任何 机器人呢
			for k,cost:=1,0; k<=f[i][1] && j-k+1>=0; k++ {
				cost += abs(f[i][0], robot[j-k+1])
				dp[i+1][j+1] =min(dp[i+1][j+1], dp[i][j-k+1] + cost)
			}
		}
	}
	return int64(dp[n][m])
}

func abs (a, b int)int {
	if a > b {
		return a - b
	}
	return b-a
}


