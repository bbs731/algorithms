package weekly

import (
	"math/bits"
)

/***

给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于 limit 。

如果不存在满足条件的子数组，则返回 0 。


示例 1：

输入：nums = [8,2,4,7], limit = 4
输出：2
解释：所有子数组如下：
[8] 最大绝对差 |8-8| = 0 <= 4.
[8,2] 最大绝对差 |8-2| = 6 > 4.
[8,2,4] 最大绝对差 |8-2| = 6 > 4.
[8,2,4,7] 最大绝对差 |8-2| = 6 > 4.
[2] 最大绝对差 |2-2| = 0 <= 4.
[2,4] 最大绝对差 |2-4| = 2 <= 4.
[2,4,7] 最大绝对差 |2-7| = 5 > 4.
[4] 最大绝对差 |4-4| = 0 <= 4.
[4,7] 最大绝对差 |4-7| = 3 <= 4.
[7] 最大绝对差 |7-7| = 0 <= 4.
因此，满足题意的最长子数组的长度为 2 。
示例 2：

输入：nums = [10,1,2,4,7,2], limit = 5
输出：4
解释：满足题意的最长子数组是 [2,4,7,2]，其最大绝对差 |2-7| = 5 <= 5 。
示例 3：

输入：nums = [4,2,2,2,4,4,2,2], limit = 0
输出：3

 */

/***
1. 思路， 维护一个 sliding window 向右移动，
	a. 如果节点 ith 可以加入 sliding window , 加入， update result
   b. 如果不可以， 移动左端点，直到合法位置。
条件， 需要直到 sliding window 内所有元素的，最小值和最大值。 就能判断 ith number 是否合法。
所以需要，维护区间内的极值， 因为不涉及到修改只是查询， 所以 ST  table 就可以。 （BIT， SegmentTree 都可以）

时间复杂度： sliding window O(n), 查询区间 O(1), 但是预处理时间 O（NlogN)  所以总复杂度应该是 O(NlogN)

不超纲的话，可以用单调队列实现， minQ 和 maxQ 这个复习的时候，可以做一下

 */

func longestSubarray(nums []int, limit int) int {
	st := NewST(nums)
	left := 0
	ans := 1

	for i := 1; i < len(nums); i++ {
		ml, mh := st.QueryMin(left, i+1), st.QueryMax(left, i+1)

		for abs(ml, nums[i]) > limit || abs(mh, nums[i]) > limit {
			left++
			ml, mh = st.QueryMin(left, i+1), st.QueryMax(left, i+1)
		}
		ans = max(ans, i-left+1)
	}
	return ans
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

//type ST [][]int
type ST struct {
	h [][]int // max
	l [][]int // min
}

func NewST(nums []int) *ST {
	n := len(nums)
	sz := bits.Len(uint(n))
	st := &ST{}
	st.h = make([][]int, n)
	st.l = make([][]int, n)
	//st := make(ST, n)
	// 初始化
	for i, v := range nums {
		st.h[i] = make([]int, sz)
		st.h[i][0] = v

		st.l[i] = make([]int, sz)
		st.l[i][0] = v
	}

	// 建表，复杂度 O(n*logn)
	for j := 1; 1<<j <= n; j++ { // j loop 的上界比较难想  0 + 1<<j <=n
		for i := 0; i+1<<j <= n; i++ {
			st.h[i][j] = max(st.h[i][j-1], st.h[i+1<<(j-1)][j-1])
			st.l[i][j] = min(st.l[i][j-1], st.l[i+1<<(j-1)][j-1])
		}
	}
	return st
}

// query  [l, r) 区间,  l, r 的下标从 0 开始   0<=l < r <=n
func (st *ST) QueryMax(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	//return st.Op(st[l][k],  st[l+1<<k][k])  这个是不对的，查询的是 [l, l+ 1<<(k+1)],右界 l + 1 <<(k+1) 可能已经超过  r 了。
	return max(st.h[l][k], st.h[r-1<<k][k])
}
func (st *ST) QueryMin(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	return min(st.l[l][k], st.l[r-1<<k][k])
}

// min, max, gcd, ...
func (*ST) Op(int, int) (_ int) { return }
