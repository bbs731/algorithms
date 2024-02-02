package dp

import (
	"sort"
)

/***

你打算利用空闲时间来做兼职工作赚些零花钱。

这里有 n 份兼职工作，每份工作预计从 startTime[i] 开始到 endTime[i] 结束，报酬为 profit[i]。

给你一份兼职工作表，包含开始时间 startTime，结束时间 endTime 和预计报酬 profit 三个数组，请你计算并返回可以获得的最大报酬。

注意，时间上出现重叠的 2 份工作不能同时进行。

如果你选择的工作在时间 X 结束，那么你可以立刻进行在时间 X 开始的下一份工作。


输入：startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]
输出：120
解释：
我们选出第 1 份和第 4 份工作，
时间范围是 [1-3]+[3-6]，共获得报酬 120 = 50 + 70。


1 <= startTime.length == endTime.length == profit.length <= 5 * 10^4
1 <= startTime[i] < endTime[i] <= 10^9
1 <= profit[i] <= 10^4

 */

/***
上面的模板代码，还是太长了， 尝试缩短一下（按照灵神关于 DP + 二分的代码）
主要是，看看是否能简化掉 ends 数组的逻辑， ends 数组应该是不需要的。
*/
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	type record struct {
		start, end, value int
	}
	records := make([]record, n+1)
	for i := 0; i < n; i++ {
		start, end, value := startTime[i], endTime[i], profit[i]
		records[i] = record{start, end, value}
	}
	// dummy node
	records[n] = record{0, 0, 0}

	sort.Slice(records, func(i, j int) bool {
		return records[i].end < records[j].end
	})

	dp := make([]int, n+1) // 这里的 dp 相当于把 end 的值域离散化了，当 end 的值域是 1e9 也没关系，只对 end 出现过的次数哟关系 n
	for i, r := range records {
		if i == 0 {
			continue // 或者让 index i 从 1 开始也可以。
		}
		// 需要找打第一个  end  <= start  的 end   >= (start+1) - 1 的位置。
		j := sort.Search(i, func(j int) bool { return records[j].end >= (r.start + 1) }) - 1
		dp[i] = max(dp[i-1], dp[j]+r.value)
	}
	return dp[n]
}

/***
还是使用了 2054 上用的模板。
DP + 二分的模板

看来这个模板是通用的， 一劳永逸，解决， 1235， 2054， 2830，2008 等 interval 可能有重叠的，线性DP 问题。
 */
func jobScheduling(startTime []int, endTime []int, profit []int) int {

	n := len(startTime)
	ends := make([]int, 0, n)
	type pair struct {
		start, value int
	}
	records := make(map[int][]pair, n)
	for i := 0; i < n; i++ {
		start, end, value := startTime[i], endTime[i], profit[i]
		records[end] = append(records[end], pair{start, value})
	}
	for k := range records {
		ends = append(ends, k)
	}
	// insert a dummy node
	ends = append(ends, 0)
	sort.Ints(ends)

	m := len(ends)
	dp := make(map[int]int, m)

	// endsIndex 的作用，就相当于 离散化， 因为在 2054的题目里， 值域太大， 离散化有效的减少了值域。
	endsIndex := make(map[int]int, len(ends))
	for i, end := range ends {
		endsIndex[end] = i
	}

	for i := 1; i < len(ends); i++ {
		end := ends[i]
		dp[end] = max(dp[end], dp[ends[endsIndex[end]-1]])
		for _, r := range records[end] {
			//dp[end] = max(dp[end], r.value)
			// 找到 第一个 end <= start  // 二分查找的技巧 <=  等价于  (>=(x+1))-1
			prevIndex := sort.SearchInts(ends, r.start+1) - 1
			prev := ends[prevIndex]
			dp[end] = max(dp[end], dp[prev]+r.value)
		}
	}
	return dp[ends[m-1]]
}
