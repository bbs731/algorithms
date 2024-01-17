package weekly

/*

给你一个整数数组 nums 和两个正整数 m 和 k 。

请你返回 nums 中长度为 k 的 几乎唯一 子数组的 最大和 ，如果不存在几乎唯一子数组，请你返回 0 。

如果 nums 的一个子数组有至少 m 个互不相同的元素，我们称它是 几乎唯一 子数组。

子数组指的是一个数组中一段连续 非空 的元素序列。



示例 1：

输入：nums = [2,6,7,3,1,7], m = 3, k = 4
输出：18
解释：总共有 3 个长度为 k = 4 的几乎唯一子数组。分别为 [2, 6, 7, 3] ，[6, 7, 3, 1] 和 [7, 3, 1, 7] 。这些子数组中，和最大的是 [2, 6, 7, 3] ，和为 18 。
示例 2：

输入：nums = [5,9,9,2,4,5,4], m = 1, k = 3
输出：23
解释：总共有 5 个长度为 k = 3 的几乎唯一子数组。分别为 [5, 9, 9] ，[9, 9, 2] ，[9, 2, 4] ，[2, 4, 5] 和 [4, 5, 4] 。这些子数组中，和最大的是 [5, 9, 9] ，和为 23 。
示例 3：

输入：nums = [1,2,1,2,1,2,1], m = 3, k = 3
输出：0
解释：输入数组中不存在长度为 k = 3 的子数组含有至少  m = 3 个互不相同元素的子数组。所以不存在几乎唯一子数组，最大和为 0 。

 */
// https://leetcode.cn/problems/maximum-sum-of-almost-unique-subarray/solutions/2424847/hua-dong-chuang-kou-fu-ti-dan-pythonjava-2vd6/

func maxSum(nums []int, m int, k int) int64 {
	ans := 0
	sum := 0
	n := len(nums)
	cnts := make(map[int]int) // 值域是 10^9 只能用 map 了吧？ 还有其它的选择吗？
	for i := 0; i < k-1; i++ {
		cnts[nums[i]]++
		sum += nums[i]
	}
	l := 0

	// 枚举右端点
	for i := k - 1; i < n; i++ {
		cnts[nums[i]]++
		sum += nums[i]
		if len(cnts) >= m {
			ans = max(ans, sum)
		}
		//从windows 中移除最左边的元素。
		sum -= nums[l]
		cnts[nums[l]]--
		if cnts[nums[l]] == 0 {
			delete(cnts, nums[l])
		}
		l++
	}
	return int64(ans)
}

// 终于遇到了灵神的官方题解
func maxSum(nums []int, m, k int) (ans int64) {
	sum := int64(0)
	cnt := map[int]int{}
	for _, x := range nums[:k-1] { // 先统计 k-1 个数
		sum += int64(x)
		cnt[x]++
	}
	for i, in := range nums[k-1:] {
		sum += int64(in) // 再添加一个数就是 k 个数了
		cnt[in]++
		if len(cnt) >= m && sum > ans {
			ans = sum
		}

		out := nums[i]
		sum -= int64(out) // 下一个子数组不包含 out，移出窗口
		cnt[out]--
		if cnt[out] == 0 {
			delete(cnt, out)
		}
	}
	return
}
