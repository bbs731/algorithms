package binary_search

import "sort"

/****
沿街有一排连续的房屋。每间房屋内都藏有一定的现金。现在有一位小偷计划从这些房屋中窃取现金。

由于相邻的房屋装有相互连通的防盗系统，所以小偷 不会窃取相邻的房屋 。

小偷的 窃取能力 定义为他在窃取过程中能从单间房屋中窃取的 最大金额 。

给你一个整数数组 nums 表示每间房屋存放的现金金额。形式上，从左起第 i 间房屋中放有 nums[i] 美元。

另给你一个整数 k ，表示窃贼将会窃取的 最少 房屋数。小偷总能窃取至少 k 间房屋。

返回小偷的 最小 窃取能力。


示例 1：

输入：nums = [2,3,5,9], k = 2
输出：5
解释：
小偷窃取至少 2 间房屋，共有 3 种方式：
- 窃取下标 0 和 2 处的房屋，窃取能力为 max(nums[0], nums[2]) = 5 。
- 窃取下标 0 和 3 处的房屋，窃取能力为 max(nums[0], nums[3]) = 9 。
- 窃取下标 1 和 3 处的房屋，窃取能力为 max(nums[1], nums[3]) = 9 。
因此，返回 min(5, 9, 9) = 5 。
示例 2：

输入：nums = [2,7,9,3,1], k = 2
输出：2
解释：共有 7 种窃取方式。窃取能力最小的情况所对应的方式是窃取下标 0 和 4 处的房屋。返回 max(nums[0], nums[4]) = 2 。


提示：

1 <= nums.length <= 10^5
1 <= nums[i] <= 10^9
1 <= k <= (nums.length + 1)/2

 */

/***
这个题的难点，就变成了， 猜一个 ans, 然后排除 所有不满足的房间，之后， 如何判断 是否满足足够的 K个房间。

思路就是：  xxx OOO xxx O x OO xxx
查一下 所有的 O 的段中满足的的最大房间数。  (len(O)+1)/2 都累加起来 判断是否 满足 >= k


牛逼啊，一次过啊！ 难度系数：2081
 */
func minCapability(nums []int, k int) int {
	l, r := 0, int(1e9)+1
	n := len(nums)

	// 这是一个 先 false, 后 true 的序列。
	for l+1 < r {
		mid := (l + r) >> 1
		start := 0
		tot := 0
		i := 0
		for i < n {
			for start < n && nums[start] > mid {
				start++
			}
			i = start
			for i < n && nums[i] <= mid {
				i++
			}
			// now reach end of the interval
			tot += (i - start + 1) / 2
			start = i
		}

		if tot >= k {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

func minCapability(nums []int, k int) int {
	//l, r := 0, int(1e9)+1
	n := len(nums)

	// 正常的 先 false, 后 true 的序列。 值域也是正常的 [0, 1e9]
	return sort.Search(int(1e9)+1, func(mid int) bool {
		start := 0
		tot := 0
		i := 0
		for i < n {
			for start < n && nums[start] > mid {
				start++
			}
			i = start
			for i < n && nums[i] <= mid {
				i++
			}
			// now reach end of the interval
			tot += (i - start + 1) / 2
			start = i
		}
		return tot >= k
	})
}

/***

https://leetcode.cn/problems/house-robber-iv/solutions/2093952/er-fen-da-an-dp-by-endlesscheng-m558/
灵神的题解：

用 DP 的思想去解， 太牛了。
f[i] 表示从从 nums[0] 到 nums[i] 中偷金额不超过  mid 的房屋数。 最后要判断  f[n-1] >=k  即， mid 是满足条件的解。

DP 用选和不选来思考。
f[i] = f[i-1]  //不选 i
f[i] = f[i-2] + 1 // 选 i

f[i] = max(f[i-1], f[i-2] + 1)
这个太牛了， 这个 DP 的时间复杂度是 O(n)


用两个变量，滚动计算！

 */

func minCapability(nums []int, k int) int {
	return sort.Search(1e9, func(mx int) bool {
		f0, f1 := 0, 0
		for _, x := range nums {
			if x <= mx {
				f0, f1 = f1, max(f1, f0+1)
			} else {
				f0 = f1
			}
		}
		return f1 >= k
	})
}
