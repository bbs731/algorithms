package one_day_exercise

import "sort"

/****
这是最朴素的想法， 但是会超时。


这个朴素的想法，说明了一点， 就是， 需要按照 endtime 来做排序。 如果不排序，
下面 loop 用 dp[c] 去更新 dp[i] 的时候，是不对的。 计算DP 的时候， 计算顺序非常重要， 这道题，就是非常好的说明。
需要 dp[c] 已经是计算好的最大值之后，才能去更新 dp[i] 这要求，我们更新 dp 的顺序，应该按照 endtime 的升序来更新。

 */

func maxValue(events [][]int, k int) int {
	n := len(events)

	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			// 不选 ith meeting, 和单选 ith
			dp[i+1][j+1] = max(dp[i][j+1], events[i][2])
			for c := 0; c < n; c++ {
				if events[i][0] > events[c][1] {
					dp[i+1][j+1] = max(dp[i+1][j+1], dp[c+1][j]+events[i][2])
				}
			}
		}
	}

	return dp[n][k]
}
