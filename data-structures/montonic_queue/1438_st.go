package montonic_queue

import "math/bits"

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

// query  [l, r) 区间  l, r 的下标从 0 开始   0<=l < r <=n
func (st *ST) QueryMax(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	//return st.Op(st[l][k],  st[l+1<<k][k])  这个是不对的，查询的是 [l, l+ 1<<(k+1)],右界 l + 1 <<(k+1) 可能已经超过  r 了。
	return max(st.h[l][k], st.h[r-1<<k][k])
}
func (st *ST) QueryMin(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	return min(st.l[l][k], st.l[r-1<<k][k])
}
