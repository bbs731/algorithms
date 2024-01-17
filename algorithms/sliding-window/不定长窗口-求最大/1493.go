package weekly

/*

给你一个二进制数组 nums ，你需要从中删掉一个元素。

请你在删掉元素的结果数组中，返回最长的且只包含 1 的非空子数组的长度。

如果不存在这样的子数组，请返回 0 。



提示 1：

输入：nums = [1,1,0,1]
输出：3
解释：删掉位置 2 的数后，[1,1,1] 包含 3 个 1 。
示例 2：

输入：nums = [0,1,1,1,0,1,1,0,1]
输出：5
解释：删掉位置 4 的数字后，[0,1,1,1,1,1,0,1] 的最长全 1 子数组为 [1,1,1,1,1] 。
示例 3：

输入：nums = [1,1,1]
输出：2
解释：你必须要删除一个元素。

 */

// 好题啊， 关键，面试的时候，就应该是这种难度的过家家
func longestSubarray(nums []int) int {
	n := len(nums)
	allones := true
	ans := 0
	prev1s := 0
	cnt1s := nums[0]

	for i := 1; i < n; i++ {
		if nums[i] == 0 {
			ans = max(ans, cnt1s+prev1s)
			if nums[i-1] == 0 {
				prev1s = 0
			} else {
				prev1s = cnt1s
			}
			cnt1s = 0
			allones = false
		} else {
			cnt1s++
			ans = max(ans, cnt1s+prev1s)
		}
	}
	if allones {
		return n - 1
	}
	return ans
}

func longestSubarray(nums []int) int {
	sum := 0
	l := 0
	r := 0
	for r = 0; r < len(nums); r++ {
		sum += nums[r]
		if sum < r-l { // 注意，这里用到的是 r-l  不是 r-l+1 而且还是 < 号。 这个技巧性太强了！无用的聪明，迁移不到其它题目上。
			sum -= nums[l]
			l++
		}
	}
	return r - l - 1
}
