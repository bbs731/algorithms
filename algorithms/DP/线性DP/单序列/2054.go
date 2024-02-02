package dp

import "sort"

/***

给你一个下标从 0 开始的二维整数数组 events ，其中 events[i] = [startTimei, endTimei, valuei] 。第 i 个活动开始于 startTimei ，结束于 endTimei ，如果你参加这个活动，那么你可以得到价值 valuei 。你 最多 可以参加 两个时间不重叠 活动，使得它们的价值之和 最大 。

请你返回价值之和的 最大值 。

注意，活动的开始时间和结束时间是 包括 在活动时间内的，也就是说，你不能参加两个活动且它们之一的开始时间等于另一个活动的结束时间。更具体的，如果你参加一个活动，且结束时间为 t ，那么下一个活动必须在 t + 1 或之后的时间开始。
输入：events = [[1,3,2],[4,5,2],[2,4,3]]
输出：4
解释：选择绿色的活动 0 和 1 ，价值之和为 2 + 2 = 4 。



输入：events = [[1,5,3],[1,5,1],[6,6,5]]
输出：8
解释：选择活动 0 和 2 ，价值之和为 3 + 5 = 8 。


2 <= events.length <= 10^5
events[i].length == 3
1 <= startTimei <= endTimei <= 10^9
1 <= valuei <= 10^6

 */

/***
这道题的值域太大了， 不能直接loop 值域。
如果DP的状态， 还是定义在值域上面， 那么就需要，优化。

dp[end] 表示， 以 end值结尾的 1 次meeting 的最大值。
用 end(i) 表示 第 ith meeting 的 end time
dp[end(i)] = max(dp[end(j), dp[end(i)]  for all j that  end(j) < end(i)

所有的 meeting 的 end time 排序得到
end(i), end(j), end(k) .....
那么对应的
dp[end(i)], dp[end(j)], dp[end(k)].... 也是升序的。


所求的答案：  需要 loop 没个 end
ans = max (ans,  dp[end_prev(i)] + value(i))

因为，可能有多个 meeting end time 相同， 所以， 需要用列表保存。loop  meetings with same end time
ans = max (ans,  dp[end_prev(i)] + value(i))   // 其中 end_prev(i) 是只， 比 start(i) 小的那个 end time 的值。因为 ends 是排序的，可以用二分来查找。


我操， 这也太难想到了， 真是锻炼思维啊。(多锻炼一下吧）
写的像屎一样， 找找灵神的答案。
https://leetcode.cn/problems/two-best-non-overlapping-events/solutions/1075386/yong-you-xian-dui-lie-wei-hu-ling-yi-ge-8ld3x/
灵神有，单调队列，简单的解法，算是特例把。 就不在 DP 的题目上使用了，现在在锻炼DP 的题目。


 */
func maxTwoEvents(events [][]int) int {
	n := len(events)
	dp := make(map[int]int, n)
	ends := make([]int, n)
	type pair struct {
		start, value int
	}
	records := make(map[int][]pair, n)
	for _, e := range events {
		start, end, value := e[0], e[1], e[2]
		records[end] = append(records[end], pair{start, value})
	}
	for k := range records {
		ends = append(ends, k)
	}
	// insert a dummy node
	ends = append(ends, 0)
	sort.Ints(ends)
	endsIndex := make(map[int]int, len(ends))
	for i, end := range ends {
		endsIndex[end] = i
	}

	ans := 0
	for _, end := range ends {
		//start, end, value := e[0], e[1], e[2]
		dp[end] = max(dp[end], dp[ends[endsIndex[end]-1]])
		for _, r := range records[end] {
			dp[end] = max(dp[end], r.value)
			// 找到 第一个 end < start  // 二分查找的技巧 <  等价于  (>=x)-1
			prevIndex := sort.SearchInts(ends, r.start) - 1
			prev := ends[prevIndex]
			//for _, r := range records[prev] {
			ans = max(ans, dp[prev]+r.value)
		}
	}
	return ans
}
