package dp

import "sort"

/***

给你一个 events 数组，其中 events[i] = [startDayi, endDayi, valuei] ，表示第 i 个会议在 startDayi 天开始，第 endDayi 天结束，如果你参加这个会议，你能得到价值 valuei 。同时给你一个整数 k 表示你能参加的最多会议数目。

你同一时间只能参加一个会议。如果你选择参加某个会议，那么你必须 完整 地参加完这个会议。会议结束日期是包含在会议内的，也就是说你不能同时参加一个开始日期与另一个结束日期相同的两个会议。

请你返回能得到的会议价值 最大和 。

 */

/**
2054 的升级题目

我靠， 感觉好爽了， 套用了模板。 然后， 把 dp 升级到二维

dp[i][j] = max (dp[i-1][j],   dp[pos][j-1] 其中 pos 是最大的 pos 满足  end(pos) < start(i) 的 pos
 */

func maxValue(events [][]int, k int) int {
	n := len(events)
	type record struct {
		start, end, value int
	}
	records := make([]record, n+1)
	for i := 0; i < n; i++ {
		start, end, value := events[i][0], events[i][1], events[i][2]
		records[i] = record{start, end, value}
	}
	records[n] = record{0, 0, 0}
	sort.Slice(records, func(i, j int) bool {
		return records[i].end < records[j].end
	})

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	// 我想知道， loop k 和 loop records 的顺序，有什么影响？ 对于这道题来说，都可以吗？
	// 测试了，本道题，怎么 loop 都可以？ 但是思考一下？ 什么情况下，是不行的？
	for j := 1; j <= k; j++ {
		for i := 1; i < len(records); i++ {
			r := records[i]
			dp[i][j] = max(dp[i-1][j], r.value) // 注意这里定义的 dp[i] 是到 i th record 为止， 只含有一个 interval 的能达到的最大值。
			// 找到 第一个 end < start  // 二分查找的技巧 <  等价于  (>=x)-1
			pos := sort.Search(i, func(k int) bool { return records[k].end >= r.start }) - 1
			//dp[i+1] = max(dp[i], dp[j]+r.value)
			dp[i][j] = max(dp[i][j], dp[pos][j-1]+r.value)
		}
	}
	return dp[n][k]
}
