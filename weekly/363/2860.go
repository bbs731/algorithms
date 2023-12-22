package weekly

import "sort"

/*


2860. 让所有学生保持开心的分组方法数

给你一个下标从 0 开始、长度为 n 的整数数组 nums ，其中 n 是班级中学生的总数。班主任希望能够在让所有学生保持开心的情况下选出一组学生：

如果能够满足下述两个条件之一，则认为第 i 位学生将会保持开心：

这位学生被选中，并且被选中的学生人数 严格大于 nums[i] 。
这位学生没有被选中，并且被选中的学生人数 严格小于 nums[i] 。
返回能够满足让所有学生保持开心的分组方法的数目。



示例 1：

输入：nums = [1,1]
输出：2
解释：
有两种可行的方法：
班主任没有选中学生。
班主任选中所有学生形成一组。
如果班主任仅选中一个学生来完成分组，那么两个学生都无法保持开心。因此，仅存在两种可行的方法。


示例 2：

输入：nums = [6,0,3,3,6,7,2,7]
输出：3
解释：
存在三种可行的方法：
班主任选中下标为 1 的学生形成一组。
班主任选中下标为 1、2、3、6 的学生形成一组。
班主任选中所有学生形成一组。

 */

//
// Constraints:
//
//
// 1 <= nums.length <= 105
// 0 <= nums[i] < nums.length

func countWays(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := 0

	if nums[0] > 0 {
		ans++
	}

	// i 是选中的人数
	for i := 1; i < n; i++ {
		if i > nums[i-1] && i < nums[i] {
			ans++
		}
	}
	if nums[n-1] < n {
		ans++
	}
	return ans
}

/*

错误的做法：
这是个 DP 的问题吗？ 用 dfs 来剪枝吗？ 看这 n =100 感觉就不行， 而且没办法加 cache
 */
func countWays(nums []int) int {

	n := len(nums)
	ans := 0
	var dfs func(int, int, int, int)
	dfs = func(i, choosed, low, high int) {
		// 剪枝
		if low >= high {
			return
		}
		// low 是递增的， high 是递减的
		if choosed <= low || choosed >= high {
			return
		}

		if i == n {
			if choosed > low && choosed < high {
				ans++
			}
			return
		}

		// 不选  selected < nums[i]
		dfs(i+1, choosed, low, min(high, nums[i]))

		//选  selected > nums[i]
		dfs(i+1, choosed+1, max(low, nums[i]), high)
	}
	dfs(0, 0, -1, n+1)
	return ans
}
