package _600_1900

import "sort"

func numFactoredBinaryTrees(arr []int) int {
	mod := int(1e9 + 7)
	n := len(arr)
	sort.Ints(arr)
	dp := make([]int, n)
	h := make(map[int]int, n)
	for i, x := range arr {
		h[x]=i
		dp[i] = 1
	}
	ans :=1
	for i:=1; i<n; i++ {
		for j:=0; j<=i-1;j++ {
			a := arr[j]
			b := arr[i]/a
			bi, ok := h[b]
			if a * b == arr[i] && ok {
				dp[i] = (dp[i] + dp[j]*dp[bi])%mod
				// loop from 0 to i-1 , a != b 自然会被 + 两遍
				//if a != b {
				//	dp[i] = (dp[i] + dp[j]*dp[bi])%mod
				//}
			}
		}
		ans = (ans+dp[i])%mod
	}
	return ans
}
