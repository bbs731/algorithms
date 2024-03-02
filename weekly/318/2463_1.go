package weekly

import (
	"math"
	"sort"
)

/***

这个想法，很牛逼啊！ 春雷。
把 厂子 i 如果有 limit capacity, 把他拆成， limit 个独立的厂子。 这样，下面我们就不需要 loop limit 了。

像灵神给的状态转移，方程：
f[i][j] 代表： 代表前 i 个工厂， 修理了 j 的机器人的最小值。
这个 DP， 就变成了选还是不选的问题了。

	f[i][j] = min(f[i-1][j], f[i-1][j-1] + abs(robot[j], facs[i]))
	f[i+1][j+1] = min(f[i][j+1], f[i][j] + abs(robot[j], facs[i]))

这个DP 的难点，在于， j 是需要反着枚举的。

return
	f[n][m]
 */
func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	facs := make([]int, 0)
	for _, f := range factory {
		pos, limit := f[0], f[1]
		for i:=1; i<=limit; i++{
			facs = append(facs, pos)
		}
	}
	sort.Ints(facs)
	m :=len(robot)
	n := len(facs)

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
			dp[i+1][j+1] = min(dp[i][j+1], dp[i][j] + abs(robot[j], facs[i]))
			//for k,cost:=1,0; k<=f[i][1] && j-k+1>=0; k++ {
			//	cost += abs(f[i][0], robot[j-k+1])
			//	dp[i+1][j+1] =min(dp[i+1][j+1], dp[i][j-k+1] + cost)
			//}
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




