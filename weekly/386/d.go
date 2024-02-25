package weekly

import "sort"

/****

这个 DP 的分析，太强了！

TsReaper 高手啊！
https://leetcode.cn/problems/earliest-second-to-mark-indices-ii/solutions/2653049/dp-by-tsreaper-3ks9/

 */
func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	n,m := len(nums), len(changeIndices)
	mp := make(map[int]int, n)
	for i:=0; i<m; i++ {
		c := changeIndices[i]-1
		if mp[c] == 0 {
			mp[c] = i+1  // 最早的可以清零的时刻。
		}
	}
	type pair struct {
		time int
		i int
	}
	l := make([]pair, 0)
	for i, v :=range mp{
		l = append(l, pair{v,  i})
	}
	sort.Slice(l, func(i, j int) bool {
		return l[i].time < l[j].time
	})

	inf := int(1e18)
	w := len(l)
	// dp[i][r]  i 时刻， r 时间上限制
	dp := make([][]int, w+1)
	for i:= range dp {
		dp[i]= make([]int, m+1)
	}
	for j:=0; j<=m; j++ {
		dp[0][j] = inf
	}
	dp[0][0] = 0

	for i:=0; i<w; i++ {
		p := l[i]
		for R:=0; R<=m; R++ {
			dp[i+1][R] = inf
		}

		for R:=0; R<=m; R++ {
			// i+1 时刻 不清零
			dp[i+1][R] = min(dp[i+1][R], dp[i][R] + nums[p.i] + 1)

			// i+1 时刻清零。
			RR := R
			if RR < p.time {
				RR = p.time+ 1
			}else {
				RR += 2
			}
			if RR > m {
				continue
			}
			dp[i+1][RR] = min(dp[i+1][RR], dp[i][R] + 2)
		}
	}

	sm := 0
	for i:=0; i<n; i++ {
		if mp[i] == 0 {
			sm += nums[i] + 1
		}
	}

	ans := inf

	for R:=0; R<=m; R++ {
		ans = min(ans, max(R, dp[w][R]+sm))
	}

	if ans >m {
		return -1
	}
	return ans
}
