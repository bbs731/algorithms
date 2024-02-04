package difference

/***
给你一个下标从 0 开始的整数数组 nums 和一个正整数 k 。

你可以对数组执行下述操作 任意次 ：

从数组中选出长度为 k 的 任一 子数组，并将子数组中每个元素都 减去 1 。
如果你可以使数组中的所有元素都等于 0 ，返回  true ；否则，返回 false 。

子数组 是数组中的一个非空连续元素序列。



示例 1：

输入：nums = [2,2,3,1,1,0], k = 3
输出：true
解释：可以执行下述操作：
- 选出子数组 [2,2,3] ，执行操作后，数组变为 nums = [1,1,2,1,1,0] 。
- 选出子数组 [2,1,1] ，执行操作后，数组变为 nums = [1,1,1,0,0,0] 。
- 选出子数组 [1,1,1] ，执行操作后，数组变为 nums = [0,0,0,0,0,0] 。
示例 2：

输入：nums = [1,3,1,1], k = 2
输出：false
解释：无法使数组中的所有元素等于 0 。


提示：

1 <= k <= nums.length <= 10^5
0 <= nums[i] <= 10^6
 */

/***
好题啊！好的面试题目。
如果实现不知道用差分数组来做的话？ 怎么办? 现场能想得到吗？


这题太tmd 难了。 折腾了，快一个小时了。其实第一遍写的就是对的， 最主要，关键的可能遗漏的点是
1. loop 的时候，应该 Loop 到  n-k 的位置. 关系到第二点
2. 因为题目只规定了，能减 不能 加， 所以恢复完数组的 diffs[i] < 0  是不合法的，需要体检返回 false。这里耗时了好长时间排除，哎！


https://leetcode.cn/problems/apply-operations-to-make-all-array-elements-equal-to-zero/solutions/2336744/chai-fen-shu-zu-pythonjavacgojs-by-endle-8qrt/
灵神的答案 ，更有启发性质， loop 不用拆成两段， 合成一段， 判断就可以。
 */
func checkArray(nums []int, k int) bool {
	n := len(nums)

	// 终于有题目， 需要预处理， 等到差分数组了。
	diffs := make([]int, n)
	diffs[0] = nums[0]
	for i := 1; i < n; i++ {
		diffs[i] = nums[i] - nums[i-1]
	}

	// 重复使用 diffs 数组, 其实是可以的。

	for i := 0; i <= n-k; i++ {
		if i > 0 {
			diffs[i] += diffs[i-1]
		}

		if diffs[i] < 0 {
			return false
		}
		// 顺序很重要
		if diffs[i] != 0 {
			if i+k < n {
				diffs[i+k] += diffs[i]
			}
		}
		diffs[i] -= diffs[i]
	}

	for i := n - k; i < n; i++ {
		if i-1 >= 0 {
			diffs[i] += diffs[i-1]
		}
		if diffs[i] != 0 {
			return false
		}
	}
	return true
}

func checkArray(nums []int, k int) bool {
	n := len(nums)

	// 终于有题目， 需要预处理， 等到差分数组了。
	diffs := make([]int, n)
	diffs[0] = nums[0]
	for i := 1; i < n; i++ {
		diffs[i] = nums[i] - nums[i-1]
	}

	// 重复使用 diffs 数组其实是可以的。 就是比较难想啊！

	for i := 0; i < n; i++ {
		if i > 0 {
			diffs[i] += diffs[i-1]
		}

		if diffs[i] == 0 {
			continue
		}

		if diffs[i] < 0 || i+k > n { // 过了 n-k 的位置，如果还不为0 就不能满足条件
			return false
		}

		// 顺序很重要
		if diffs[i] != 0 {
			if i+k < n {
				diffs[i+k] += diffs[i]
			}
		}
		diffs[i] -= diffs[i]

	}
	return true
}
