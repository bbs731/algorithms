package dp

/***
给你两个数组 nums1 和 nums2 。

请你返回 nums1 和 nums2 中两个长度相同的 非空 子序列的最大点积。

数组的非空子序列是通过删除原数组中某些元素（可能一个也不删除）后剩余数字组成的序列，但不能改变数字间相对顺序。比方说，[2,3,5] 是 [1,2,3,4,5] 的一个子序列而 [1,5,3] 不是。


示例 1：
输入：nums1 = [2,1,-2,5], nums2 = [3,0,-6]
输出：18
解释：从 nums1 中得到子序列 [2,-2] ，从 nums2 中得到子序列 [3,-6] 。
它们的点积为 (2*3 + (-2)*(-6)) = 18 。

示例 2：
输入：nums1 = [3,-2], nums2 = [2,-6,7]
输出：21
解释：从 nums1 中得到子序列 [3] ，从 nums2 中得到子序列 [7] 。
它们的点积为 (3*7) = 21 。

示例 3：
输入：nums1 = [-1,-1], nums2 = [1,1]
输出：-1
解释：从 nums1 中得到子序列 [-1] ，从 nums2 中得到子序列 [1] 。
它们的点积为 -1 。


提示：

1 <= nums1.length, nums2.length <= 500
-1000 <= nums1[i], nums2[i] <= 100


点积：

定义 a = [a1, a2,…, an] 和 b = [b1, b2,…, bn] 的点积为：

\mathbf{a}\cdot \mathbf{b} = \sum_{i=1}^n a_ib_i = a_1b_1 + a_2b_2 + \cdots + a_nb_n

这里的 Σ 指示总和符号。
 */

/***
dp[i][j][k] =
1. dp[i-1][j-1][k-1] +nums[i]nums[j]
2. dp[i-1][j][k] , dp[i][j-1][k], dp[i-1][j-1][k]


这道题的状态， 可以不用设计 k?  直接用 dp[i][j] 为什么？
dp[i-1][j] 时候的条件是 i>0
dp[i][j-1] 时的条件是 j> 0

哎， 找个机会重新写一下。

 */

func maxDotProduct(nums1 []int, nums2 []int) int {
	inf := int(1e15)
	m, n := len(nums1), len(nums2)
	p := min(m, n)
	f := make([][][]int, m+1)
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, p+1)
		}
	}
	// 初始化， 这个是不是太难了。
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			for k := 0; k <= p; k++ {
				if k == 0 {
					f[i][j][0] = 0
				}
				if i < k || j < k { // 这个写成了 && 愚蠢啊。
					f[i][j][k] = -inf
				}
			}
		}
	}

	ans := -inf
	for k := 1; k <= p; k++ {
		for i := k - 1; i < m; i++ {
			for j := k - 1; j < n; j++ {
				f[i+1][j+1][k] = max(f[i][j][k-1]+nums1[i]*nums2[j], f[i][j+1][k], f[i+1][j][k], f[i][j][k])
				ans = max(ans, f[i+1][j+1][k])
			}
		}
	}
	return ans
}
