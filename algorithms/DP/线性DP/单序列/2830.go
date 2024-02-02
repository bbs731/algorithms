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
		//dp[o[1]] = max(dp[o[1]], o[2])
		group[end] = append(group[end], pair{start, gold})
	}

	for end, g := range group {
		//dp[end] = max(dp[end], dp[end-1])
		dp[end+1] = max(dp[end+1], dp[end])
		for _, p := range g {
			//if p.start > 0 {
			dp[end+1] = max(dp[end+1], dp[p.start]+p.gold)
			//}
		}
		//if offers[pos][1] < start {
		//	dp[offers[i][1]] = max(dp[offers[i][1]], dp[offers[pos][1]]+offers[i][2])
		//}
	}
	return dp[n]
}

func maximizeTheProfit(n int, offers [][]int) (ans int) {
	sort.Slice(offers, func(i, j int) bool {
		return offers[i][1] < offers[j][1]
	})

	dp := make([]int, n)
	m := len(offers)
	endm := make(map[int]struct{})
	ends := make([]int, 0, m)
	for _, o := range offers {
		dp[o[1]] = max(dp[o[1]], o[2])
		endm[o[1]] = struct{}{}
		//if i > 0 {
		//	// 保证 dp 是一个递增的序列
		//	dp[o[1]] = max(dp[o[1]], dp[offers[i-1][1]])
		//}
	}
	for k := range endm {
		ends = append(ends, k)
	}
	sort.Ints(ends)

	for i := 1; i < len(ends); i++ {
		dp[ends[i]] = max(dp[ends[i]], dp[ends[i-1]])
	}

	//fmt.Println(dp)

	for _, o := range offers {
		dp[o[1]] = max(dp[o[1]], ans)
		start := o[0]
		//pos := sort.Search(m, func(k int) bool {
		//	k++
		//	return offers[k][1] > start
		//})
		pos := sort.SearchInts(ends, start-1)

		if ends[pos] < start {
			dp[o[1]] = max(dp[o[1]], dp[ends[pos]]+o[2])
		}
		//if offers[pos][1] < start {
		//	dp[offers[i][1]] = max(dp[offers[i][1]], dp[offers[pos][1]]+offers[i][2])
		//}
		ans = max(ans, dp[o[1]])
	}
	return
}
