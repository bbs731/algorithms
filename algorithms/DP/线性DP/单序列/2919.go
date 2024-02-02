package dp

/***

给你一个下标从 0 开始、长度为 n 的整数数组 nums ，和一个整数 k 。

你可以执行下述 递增 运算 任意 次（可以是 0 次）：

从范围 [0, n - 1] 中选择一个下标 i ，并将 nums[i] 的值加 1 。
如果数组中任何长度 大于或等于 3 的子数组，其 最大 元素都大于或等于 k ，则认为数组是一个 美丽数组 。

以整数形式返回使数组变为 美丽数组 需要执行的 最小 递增运算数。

子数组是数组中的一个连续 非空 元素序列。


示例 1：

输入：nums = [2,3,0,0,2], k = 4
输出：3
解释：可以执行下述递增运算，使 nums 变为美丽数组：
选择下标 i = 1 ，并且将 nums[1] 的值加 1 -> [2,4,0,0,2] 。
选择下标 i = 4 ，并且将 nums[4] 的值加 1 -> [2,4,0,0,3] 。
选择下标 i = 4 ，并且将 nums[4] 的值加 1 -> [2,4,0,0,4] 。
长度大于或等于 3 的子数组为 [2,4,0], [4,0,0], [0,0,4], [2,4,0,0], [4,0,0,4], [2,4,0,0,4] 。
在所有子数组中，最大元素都等于 k = 4 ，所以 nums 现在是美丽数组。
可以证明无法用少于 3 次递增运算使 nums 变为美丽数组。
因此，答案为 3 。
示例 2：

输入：nums = [0,1,3,3], k = 5
输出：2
解释：可以执行下述递增运算，使 nums 变为美丽数组：
选择下标 i = 2 ，并且将 nums[2] 的值加 1 -> [0,1,4,3] 。
选择下标 i = 2 ，并且将 nums[2] 的值加 1 -> [0,1,5,3] 。
长度大于或等于 3 的子数组为 [0,1,5]、[1,5,3]、[0,1,5,3] 。
在所有子数组中，最大元素都等于 k = 5 ，所以 nums 现在是美丽数组。
可以证明无法用少于 2 次递增运算使 nums 变为美丽数组。
因此，答案为 2 。


3 <= n == nums.length <= 10^5
0 <= nums[i] <= 10^9
0 <= k <= 109


*/


/***

https://leetcode.cn/problems/minimum-increment-operations-to-make-array-beautiful/solutions/2503157/qiao-miao-she-ji-zhuang-tai-xuan-huo-bu-8547u/
灵神题解

https://www.bilibili.com/video/BV1tw411q7VZ/?vd_source=84c3c489cf545fafdbeb3b3a6cd6a112
有详细的， dfs 的推导过程。

 */
/***
dp[i][1] 选 i 之后 的最优解
 dp[i][0] 表示不选 i 的最优解。

 dp[i][1] = min(dp[i][1], dp[i+3][1] +  k - nums[i])
 dp[i][0] = min(dp[i+1][1] + k - nums[i-1], dp[i+2][1] + k - nums[i-2]）

 倒序枚举 i 就可以了吧。
*/


func minIncrementOperations(nums []int, k int) int64 {
	n := len(nums)

	//if n == 3 {
	//	return getCost(max(nums[0], nums[1], nums[2]), k)
	//}

	inf := int(1e16)
	dp := make([][2]int, n)
	for i := range dp {
		dp[i][0] = inf
		dp[i][1] = inf
	}
	// 初始化 n-1, n-2, n-3
	for i:=0; i<=2; i++ {
		dp[n-1-i][1] = getCost(nums[n-1-i], k)
	}
	// 没人关心 dp[n-1][0]
	for i := n - 1; i-3 >= 0; i-- {
		dp[i-2][0] = min(dp[i-2][0], dp[i][1])
		dp[i-1][0] = min(dp[i-1][0], dp[i][1])
		dp[i-3][1] = min(dp[i-3][1], dp[i][1] + getCost(nums[i-3], k))
		dp[i-3][1] = min(dp[i-3][1], dp[i-1][1] + getCost(nums[i-3], k))
		dp[i-3][1] = min(dp[i-3][1], dp[i-2][1] + getCost(nums[i-3], k))
	}

	ans := inf
	for i:=2; i>=0; i-- {
		ans = min(ans, dp[i][1])
	}
	return int64(ans)
}


/***
dp[i] 选 i 之后 的最优解
dp[i][0] 表示不选 i 的最优解。

dp[i] = min(dp[i], dp[i+3] +  k - nums[i], dp[i+2] + k-nums[i, dp[i+1] + k-nums[i])

 倒序枚举 i 就可以了吧。
*/


func getCost(a, k int) int {
	if a >= k {
		return 0
	}
	return k - a
}

func minIncrementOperations(nums []int, k int) int64 {
	n := len(nums)
	inf := int(1e16)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = inf
	}
	// 初始化 n-1, n-2, n-3
	for i:=0; i<=2; i++ {
		dp[n-1-i] = getCost(nums[n-1-i], k)
	}
	for i := n - 1; i-3 >= 0; i-- {
		dp[i-3] = min(dp[i-3], min(dp[i],dp[i-1], dp[i-2]) + getCost(nums[i-3], k))
	}

	ans := inf
	for i:=2; i>=0; i-- {
		ans = min(ans, dp[i])
	}
	return int64(ans)
}

/***
https://leetcode.cn/problems/minimum-increment-operations-to-make-array-beautiful/solutions/2503199/xiao-yang-xiao-en-dong-tai-gui-hua-zi-sh-hh9c/
按照小羊的题解，再简化。
操， 线性DP
 */
func minIncrementOperations(nums []int, k int) int64 {
	var dp1, dp2, dp3 int
	for _, num := range nums{
		dp1, dp2, dp3 = min(dp1, dp2, dp3) + max(0, k-num), dp1, dp2
	}
	return int64(min(dp1, dp2, dp3))
}
