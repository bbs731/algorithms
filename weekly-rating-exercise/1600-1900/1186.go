package _600_1900

import "math"

// 看看灵神的题解，总是那么的让人震撼！
//https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/solutions/2321829/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-hzz6/

//	​
//
//dfs(i,0)=max(dfs(i−1,0),0)+arr[i]
//dfs(i,1)=max(dfs(i−1,1)+arr[i],dfs(i−1,0))
//​
//
//f[i+1][0]=max(f[i][0],0)+arr[i]
//f[i+1][1]=max(f[i][1]+arr[i],f[i][0])
//​


func maximumSum(arr []int) int {
	ans := math.MinInt
	f := make([][2]int, len(arr)+1)
	f[0] = [2]int{math.MinInt / 2, math.MinInt / 2} // 除 2 防止负数相加溢出
	for i, x := range arr {
		f[i+1][0] = max(f[i][0], 0) + x
		f[i+1][1] = max(f[i][1]+x, f[i][0])
		ans = max(ans, max(f[i+1][0], f[i+1][1]))
	}
	return ans
}

//  自己的解法和上面有差距， 需要特判的情况还是太多。 缩小差距， 其实主要是思路上的差距， 有特殊到一般的思考方式。多训练。
//

// dp[i]  以 index i 结尾的最大子串
//  dp[i] = max(dp[i-1] + nums[i], 0)
//  dp[i+1] = max(dp[i] + nums[i], 0)

//   f[i]  以 index i 开始的最大子串值。
//   f[i] = max(f[i+1] + nums[i], 0)

func maximumSum(arr []int) int {
	n := len(arr)
	dp := make([]int, n+1)
	f := make([]int, n+1)
	allneg := true
	largest := arr[0]
	ans := 0
	for i, x := range arr {
		dp[i+1]=max(dp[i] + x, 0)
		if x >= 0 {
			allneg = false
		}
		largest = max(largest, x)
	}
	for i:=n-1; i>=0; i-- {
		f[i] = max(f[i+1] + arr[i], 0)
	}

	for i, x :=range arr {
		ans = max(ans, dp[i], f[i])
		if x < 0 {
			ans = max(ans, dp[i] + f[i+1])
		}
	}
	if allneg {
		return largest
	}
	return ans
}
