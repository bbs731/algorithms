package dp

import (
	"fmt"
	"sort"
)

/***

这道题，卡死我了。
果然，线性DP 的问题，会要命啊。
 */

func maximizeTheProfit(n int, offers [][]int) int {
	type pair struct{ start, gold int }
	groups := make([][]pair, n)
	for _, offer := range offers {
		start, end, gold := offer[0], offer[1], offer[2]
		groups[end] = append(groups[end], pair{start, gold})
	}

	f := make([]int, n+1)
	for end, g := range groups {
		f[end+1] = f[end]
		for _, p := range g {
			f[end+1] = max(f[end+1], f[p.start]+p.gold)
		}
	}
	fmt.Println(f)
	return f[n]
}

/***
留下耻辱的记忆把， 这个做不出来。
这题要了命了，做了能有12个小时。 多训练这种类型的题目。
 */

func maximizeTheProfit(n int, offers [][]int) (ans int) {

	dp := make([]int, n+1)
	type pair struct{ start, gold int }
	// 难在了，如何处理 end 相同的重复元素
	group := make([][]pair, n)
	for _, offer := range offers {
		start, end, gold := offer[0], offer[1], offer[2]
		group[end] = append(group[end], pair{start, gold})
	}

	for end, g := range group {
		dp[end+1] = max(dp[end+1], dp[end])
		for _, p := range g {
			dp[end+1] = max(dp[end+1], dp[p.start]+p.gold)
		}
	}
	return dp[n]
}

/***
用极其复杂的思路，做完了， 2054之后， 在回来考虑一下， 最初 2830 的想法，
尝试一下。


想用二分，就满足你。就是逻辑有点复杂。证明了 2054 上的方法是通用的， 可以解决
2830， 2008 这类有 interval 的问题。 [start, end] 的问题。这里面有陷阱需要考虑：
1. 就是 有多个 intervals 有相同的 end time 这时候，需要都保存起来，然后 Loop 更新 dp[end]
2. 这个下标，是难点. 因为 我们需要找前一个 end 的 index; 还有 找 < start 的二分 是 (>=start) - 1 值域都有可能是 -1
需要特判， 所以，不如把原始的 index 都加 1 然后， insert 一个 dummy 的 0 在 ends 中， 这样可以减少上面提到的两种情况 index 的检查。

 */

func maximizeTheProfit(n int, offers [][]int) (ans int) {
	m := len(offers)
	dp := make(map[int]int, n+1)
	ends := make([]int, 0, m)
	type pair struct {
		start, value int
	}
	records := make(map[int][]pair, n)
	for _, e := range offers {
		// 把 区间的左右端点 start, end 都加1 方便后面处理。
		start, end, value := e[0]+1, e[1]+1, e[2]
		records[end] = append(records[end], pair{start, value})
	}
	for k := range records {
		ends = append(ends, k)
	}
	// insert a dummy node
	ends = append(ends, 0)
	sort.Ints(ends)

	// endsIndex 的作用，就相当于 离散化， 因为在 2054的题目里， 值域太大， 离散化有效的减少了值域。
	endsIndex := make(map[int]int, len(ends))
	for i, end := range ends {
		endsIndex[end] = i
	}

	for i := 1; i < len(ends); i++ {
		end := ends[i]
		//start, end, value := e[0], e[1], e[2]
		dp[end] = max(dp[end], dp[ends[endsIndex[end]-1]])
		for _, r := range records[end] {
			dp[end] = max(dp[end], r.value)
			// 找到 第一个 end < start  // 二分查找的技巧 <  等价于  (>=x)-1
			prevIndex := sort.SearchInts(ends, r.start) - 1
			prev := ends[prevIndex]
			//for _, r := range records[prev] {
			dp[end] = max(dp[end], dp[prev]+r.value)
		}
	}
	return dp[ends[len(ends)-1]]
}

func maximizeTheProfit(n int, offers [][]int) (ans int) {
	m := len(offers)
	dp := make(map[int]int, n+1)
	ends := make([]int, 0, m)
	type pair struct {
		start, value int
	}
	records := make(map[int][]pair, n)
	for _, e := range offers {
		// 把 区间的左右端点 start, end 都加1 方便后面处理。
		start, end, value := e[0], e[1], e[2]
		records[end] = append(records[end], pair{start, value})
	}
	for k := range records {
		ends = append(ends, k)
	}
	// insert a dummy node
	//ends = append(ends, 0)
	sort.Ints(ends)
	endsIndex := make(map[int]int, len(ends))
	for i, end := range ends {
		endsIndex[end] = i
	}

	for i := 0; i < len(ends); i++ {
		end := ends[i]
		//start, end, value := e[0], e[1], e[2]
		if endsIndex[end] > 0 {
			dp[end] = max(dp[end], dp[ends[endsIndex[end]-1]])
		}
		for _, r := range records[end] {
			dp[end] = max(dp[end], r.value)
			// 找到 第一个 end < start  // 二分查找的技巧 <  等价于  (>=x)-1
			prevIndex := sort.SearchInts(ends, r.start) - 1
			if prevIndex >= 0 {
				prev := ends[prevIndex]
				//for _, r := range records[prev] {
				dp[end] = max(dp[end], dp[prev]+r.value)
			}
		}
	}
	return dp[ends[len(ends)-1]]
}
