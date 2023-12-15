package _600_1900

import (
	"fmt"
	"sort"
)

/*
// dp[i] = min(dp[i+1] + cost[0], dp[i+7] + cost[1], dp[i+30] + cost[2])
初始化条件是什么?
*/
func mincostTickets(days []int, costs []int) int {
	n := len(days)
	dp := make([]int, days[n-1]+31)
	j := n - 1

	for i := days[n-1]; i >= 1; i-- {
		if j >= 0 && i == days[j] {
			dp[i] = min(dp[i+1]+costs[0], dp[i+7]+costs[1], dp[i+30]+costs[2])
			j--
		} else {
			// 不是  days[] 里的日期，不用花费。
			dp[i] = dp[i+1]
		}
	}
	return dp[1]
}

// 不知道错在哪里了。
func mincostTickets(days []int, costs []int) int {
	days = append([]int{0}, days...)
	n := len(days)
	ans := make([]int, len(days))
	interval := []int{1, 7, 30}
	//ans[0] = min(costs[0], costs[1], costs[2])
	for i := 1; i < n; i++ {
		x := days[i]
		ans[i] = ans[i-1] + costs[0]
		// covered 1 , 7, 30 days
		for j := 1; j < 3; j++ {
			c := costs[j]
			before := x - interval[j]
			if before <= 0 {
				ans[i] = min(ans[i], ans[0]+c)
			} else {
				pos := sort.SearchInts(days[:i], before)
				if pos == i {
					ans[i] = min(ans[i], ans[i-1]+c)
				} else {
					ans[i] = min(ans[i], ans[pos]+c)
				}
			}
		}
	}
	fmt.Println(ans)
	return ans[len(ans)-1]
}
