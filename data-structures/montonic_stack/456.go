package montonic_stack

/****
给你一个整数数组 nums ，数组中共有 n 个整数。132 模式的子序列 由三个整数 nums[i]、nums[j] 和 nums[k] 组成，并同时满足：i < j < k 和 nums[i] < nums[k] < nums[j] 。

如果 nums 中存在 132 模式的子序列 ，返回 true ；否则，返回 false 。



示例 1：

输入：nums = [1,2,3,4]
输出：false
解释：序列中不存在 132 模式的子序列。
示例 2：

输入：nums = [3,1,4,2]
输出：true
解释：序列中有 1 个 132 模式的子序列： [1, 4, 2] 。
示例 3：

输入：nums = [-1,3,2,0]
输出：true
解释：序列中有 3 个 132 模式的的子序列：[-1, 3, 2]、[-1, 3, 0] 和 [-1, 2, 0] 。


提示：

n == nums.length
1 <= n <= 2 * 10^5
-10^9 <= nums[i] <= 10^9

 */

/****
这是一道很好的面试题目啊！


下面的思路，我最初的解题想法， 可能有些复杂： (去找题解）

譬如，这个序列 【 7，8， 3，4， 1， 5,  6]

处理完 7, 8,  把  {7,8} 这个区间放在 stack 上 [{7,8}]
处理完 3, 4   把  {3,4} 这个形成的区间放在 stack 上 [{7,8}, {3, 4}]
处理完 1, 5   这时发现 5 大于 stack 栈顶的这个区间的右端点， 5 > 4， 所以可以归并 {1, 5} 和 {3,4} 得到 [ min(1, 3}, 5] 放在 stack 上 [{7,8}, {1,5}]
处理 6, 看栈顶区间， 6 > 5, 合并区间 最后变成 [{7,8}, { min(1, inf}, 6}]

合并取件单额时候，顺便看看有没有所有的解。
因为所有元素， 进栈出栈， 最多 1 次， 所以可以保证， 时间复杂度是线性的。

春雷，你也是够牛的， 这么复杂的逻辑，你也能写对！ 而且基本上一遍过。


下面的这个解法是： 宫三千，纯粹利用单调栈的解法，非常简单，利用了单调栈的性质。
https://leetcode.cn/problems/132-pattern/solutions/676970/xiang-xin-ke-xue-xi-lie-xiang-jie-wei-he-95gt/


 */
type pair struct {
	left, right int
}

func find132pattern(nums []int) bool {
	inf := int(1e11)
	st := []pair{}
	l := inf

	for _, v := range nums {
		// 尝试用 v 来归并， st 上的区间
		for len(st) > 0 && v >= st[len(st)-1].right {
			// in case l already got a value before. e.g test case [40,50,25,35,15,35,20]
			l = min(st[len(st)-1].left, l)
			// pop st
			st = st[:len(st)-1]
		}
		// check whether v is the answer so call k between i,j ,k
		if len(st) > 0 && v > st[len(st)-1].left {
			return true
		}

		if l != inf && v > l {
			// append interval
			st = append(st, pair{l, v})
			// reset l to inf to form new intervals for follow processing
		} else {
			// otherwise use v as l  the left side of new interval
			l = v
		}
	}
	return false
}

/***
来吧 宫三，简单的解法

思路，如果是对的， 代码就会变简单啊。 利用了单调栈的性质， 但是这个性质好难想啊！
https://leetcode.cn/problems/132-pattern/solutions/676970/xiang-xin-ke-xue-xi-lie-xiang-jie-wei-he-95gt/
 */
func find132pattern(nums []int) bool {
	st := []int{}
	n := len(nums)
	k := -int(1e9)

	// 需要从右向左遍历
	for i := n - 1; i >= 0; i-- {
		v := nums[i]
		if k > v {
			return true
		}
		for len(st) > 0 && v > st[len(st)-1] {
			k = max(k, st[len(st)-1]) // k keep the max value popped from stack
			// pop stack
			st = st[:len(st)-1]
		}
		st = append(st, v)
	}
	return false
}
