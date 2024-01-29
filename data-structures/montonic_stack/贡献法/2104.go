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
