package weekly

/***

给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数 。



示例 1：

输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。
示例 2：

输入：nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。


提示：

1 <= nums.length <= 105
nums[i] 不是 0 就是 1
0 <= k <= nums.length

 */

func longestOnes(nums []int, k int) int {
	//n := len(nums)
	ans := 0
	l := 0
	cntZero := 0

	for i, n := range nums {
		if n == 0 {
			cntZero++
		}

		for cntZero > k {
			// move left to satisfy constrain
			if nums[l] == 0 {
				cntZero--
			}
			l++
		}
		ans = max(ans, i-l+1)
	}

	return ans
}
