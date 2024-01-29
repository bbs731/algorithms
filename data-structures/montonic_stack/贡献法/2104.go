package montonic_stack

import "math/bits"

/***

给你一个整数数组 nums 。nums 中，子数组的 范围 是子数组中最大元素和最小元素的差值。

返回 nums 中 所有 子数组范围的 和 。

子数组是数组中一个连续 非空 的元素序列。


示例 1：

输入：nums = [1,2,3]
输出：4
解释：nums 的 6 个子数组如下所示：
[1]，范围 = 最大 - 最小 = 1 - 1 = 0
[2]，范围 = 2 - 2 = 0
[3]，范围 = 3 - 3 = 0
[1,2]，范围 = 2 - 1 = 1
[2,3]，范围 = 3 - 2 = 1
[1,2,3]，范围 = 3 - 1 = 2
所有范围的和是 0 + 0 + 0 + 1 + 1 + 2 = 4
示例 2：

输入：nums = [1,3,3]
输出：4
解释：nums 的 6 个子数组如下所示：
[1]，范围 = 最大 - 最小 = 1 - 1 = 0
[3]，范围 = 3 - 3 = 0
[3]，范围 = 3 - 3 = 0
[1,3]，范围 = 3 - 1 = 2
[3,3]，范围 = 3 - 3 = 0
[1,3,3]，范围 = 3 - 1 = 2
所有范围的和是 0 + 0 + 0 + 2 + 0 + 2 = 4
示例 3：

输入：nums = [4,-2,-3,4,1]
输出：59
解释：nums 中所有子数组范围的和是 59


提示：
1 <= nums.length <= 1000
-10^9 <= nums[i] <= 10^9

 */

/***
贡献法，不知道怎么做。

先尝试手撸一遍 ST table。
做到 range query O(1) 然后  O（n^2) n =1000 的复杂度可以过。

用的 灵神 ST table 的模板：
https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/sparse_table.go#L29


ToDo:  找时间，再写 Sparse Table 的实现吧， 还是会错误，还是不熟练。 index 太难了！
 */

/***

来，尝试一下单调栈的写法：
 */

/****
这个 left, right 的模板，感觉比 907 的清楚一些。

好好， 理解 单调栈吧， 有点太难想了！

https://leetcode.cn/problems/sum-of-subarray-ranges/solutions/1153054/cong-on2-dao-ondan-diao-zhan-ji-suan-mei-o1op/
 */
func solve(nums []int) (ans int64) {
	n := len(nums)
	left := make([]int, n)  // left[i] 为左侧严格大于 num[i] 的最近元素位置（不存在时为 -1）
	right := make([]int, n) // right[i] 为右侧大于等于 num[i] 的最近元素位置（不存在时为 n）
	for i := range right {
		right[i] = n
	}
	st := []int{-1}
	for i, v := range nums {
		for len(st) > 1 && nums[st[len(st)-1]] <= v {
			right[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}
	for i, v := range nums {
		ans += (int64(i-left[i])*int64(right[i]-i) - 1) * int64(v)
	}
	return
}

func subArrayRanges(nums []int) int64 {
	ans := solve(nums)
	for i, v := range nums { // 小技巧：所有元素取反后算的就是最小值的贡献
		nums[i] = -v
	}
	return ans + solve(nums)
}

//
//func solve(nums []int) int64 {
//	st := []int{-1}
//	ans := 0
//
//	for r, x := range nums {
//		for len(st) > 1 && nums[st[len(st)-1]] <= x {
//			i := st[len(st)-1]
//			st = st[:len(st)-1]
//			ans -= nums[i] * ((i - st[len(st)-1]) * (r - i)) // nums[i] 作为 最小值的贡献
//		}
//		st = append(st, r)
//	}
//
//	return int64(ans)
//}
//
//func subArrayRanges(nums []int) int64 {
//	nums = append(nums, int(1e10))
//	ans := solve(nums)
//	for i, v := range nums { // 小技巧：所有元素取反后算的就是最小值的贡献
//		nums[i] = -v
//	}
//	nums[len(nums)-1] = int(1e10)
//	return -(ans + solve(nums))
//}

func subArrayRanges(nums []int) (ans int64) {
	for i, num := range nums {
		min, max := num, num
		for _, v := range nums[i+1:] {
			if v < min {
				min = v
			} else if v > max {
				max = v
			}
			ans += int64(max - min)
		}
	}
	return
}

/******
哎， 为啥要用 st table, 看看， 上面， 既然暴力的话， 区间中的最大，最小值， 可以在 loop 中得到， 不需要 ST query。
 */
func subArrayRanges(nums []int) int64 {
	n := len(nums)
	sz := bits.Len(uint(n))

	// bits 长度。
	//nb := bits.Len(uint(n))

	type pair struct {
		min, max int
	}

	st := make([][]pair, n)
	for i, v := range nums {
		st[i] = make([]pair, sz)
		st[i][0].min = v
		st[i][0].max = v
	}

	// build up st table
	for j := 1; 1<<j <= n; j++ { // 坑一， 需要先 loop j , 再 loop i
		for i := 0; i+1<<uint(j) <= n; i++ {
			st[i][j].min = min(st[i][j-1].min, st[i+1<<(j-1)][j-1].min)
			st[i][j].max = max(st[i][j-1].max, st[i+1<<(j-1)][j-1].max)
		}
	}

	// [l, r)   // 坑二，  query 的区间是个 [l, r)
	query := func(l, r int) (int, int) {
		k := bits.Len32(uint32(r-l)) - 1 // 这里，太难了！
		return min(st[l][k].min, st[r-1<<k][k].min), max(st[l][k].max, st[r-1<<k][k].max)
	}

	ans := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			m1, m2 := query(i, j+1) // 因为是 [l, r) 所以这里  j 需要 + 1
			ans += m2 - m1
		}
	}
	return int64(ans)
}
