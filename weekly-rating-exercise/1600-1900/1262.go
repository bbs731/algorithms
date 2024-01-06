package _600_1900

import "fmt"

// dp[n][0]
// nums[n] % 3 = 1
//  dp[n][1] = dp[n-1][1] + dp[n-1][0]+ nums[n]
//  dp[n][0] = dp[n-1][0] + dp[n-1][2] + nums[n]
//  dp[n][2] = dp[n-1][2] + dp[n-1][1] + nums[n]

//   dp[n][0] = dp[n-1][(3-nums[n]%3)%3] + nums[n]
//   dp[n][1] = dp[n-1][1]
//   dp[n][2] = dp[n-1][2]

//		 3, 6, 5, 1  8
// 0	 3, 9, 9  15  
// 1			  1
// 2           5  5

// 灵神的题解，  我差了十万八千里啊！
// 这个 dp[0][1] = dp[0][2] = -inf 这个初始化是点睛之笔啊！
// 看看， dfs 的版本是真么写的?
// https://leetcode.cn/problems/greatest-sum-divisible-by-three/solutions/2313700/liang-chong-suan-fa-tan-xin-dong-tai-gui-tsll/

func maxSumDivThree(nums []int) int {
	n := len(nums)
	dp := make([][3]int, n+1)
	//dp[0][nums[0]%3] = nums[0]
	dp[0][1] = -1e10
	dp[0][2] = -1e10
	for i:=0; i<n; i++ {
		for j:=0; j<=2; j++ {
			dp[i+1][j] = max(dp[i][j], dp[i][(3+(j-nums[i])%3)%3] + nums[i])
		}
	}
	return dp[n][0]
}


func maxSumDivThree(nums []int) int {
	n := len(nums)
	dp := make([][3]int, n+1)
	//dp[0][nums[0]%3] = nums[0]
	for i:=0; i<n; i++ {
		m := nums[i]%3
		dp[i+1][m] = dp[i][m]
		dp[i+1][(m+1)%3] = dp[i][(m+1)%3]
		dp[i+1][(m+2)%3] = dp[i][(m+2)%3]


		if m == 0 {
			dp[i+1][0] += nums[i]
			if dp[i+1][1] !=0 {
				dp[i+1][1] += nums[i]
			}
			if dp[i+1][2] != 0 {
				dp[i+1][2] += nums[i]
			}
		} else if m == 1 {
			dp[i+1][1] = max(dp[i+1][1], dp[i][0]+nums[i])
			if dp[i][2] !=0 {
				dp[i+1][0] = max(dp[i+1][0], dp[i][2]+nums[i])
			}
			if dp[i][1] !=0 {
				dp[i+1][2] = max(dp[i+1][2], dp[i][1]+nums[i])
			}
		} else {
			if dp[i][1] != 0 {
				dp[i+1][0] = max(dp[i+1][0], dp[i][1]+nums[i])
			}
			if dp[i][2] !=0 {
				dp[i+1][1] = max(dp[i+1][1], dp[i][2]+nums[i])
			}
			dp[i+1][2] = max(dp[i+1][2], dp[i][0]+nums[i])

		}
		//dp[i+1][m] = dp[i][m] + nums[i]
		//if dp[i+1][(2-m)%3] !=0 {
		//	dp[i+1][(2-m)%3] += nums[i]
		//}
		//if dp[i+1][(4-m)%3] != 0 {
		//	dp[i+1][(4-m)%3] += nums[i]
		//}
		//dp[i+1][(m+1)%3] = dp[i][(2-m)%3] + nums[i]
		//dp[i+1][(m+2)%3] = dp[i][(4-m)%3] + nums[i]
	}
	return dp[n][0]
}
