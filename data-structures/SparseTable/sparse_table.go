package SparseTable

import "math/bits"

/*
https://oi-wiki.org/ds/sparse-table/
https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/sparse_table.go

结论：
1. ST 建表的时间复杂度 O（n*logn)  查询的时间复杂度 O(1)
2. ST table 不支持修改 （需要修改看线段树 segment tree)
3. ST table 只适合 Op 是重复贡献的操作， 即 x Op x = x， 可以简单理解为， Op 是 max, min, gcd


分析： 以 Op 为 max 为例
令 f(i,j) 表示区间 [i,i+2^j-1] 的最大值。

得到状态转移方程：
f(i,j) = max(f(i, j-1), f(i + 1 <<(j-1), j-1))
初始化条件：
f(i, 0) = nums[i]

遍历顺序：
j 以正序遍历， i 因为没有降维， 所以，正序遍历还好理解。
*/

// 验证题目 LC1438
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

// 注意查询，是左闭右开的区间
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
