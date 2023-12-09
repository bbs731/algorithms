package SegmentTree

type SegmentTreeB struct {
	n  int
	dp []int
	b  []int // lazy marker，一种优化的手段. 这个只是对 sum 操作的优化， 对于 max，这种优化不适用! 所以感觉这个版本的 segmentreeB 其实不是很实用
}

func (seg *SegmentTreeB) Init(nums []int) {
	n := len(nums)
	seg.n = n
	seg.dp = make([]int, 4*n)
	seg.b = make([]int, 4*n)

	var build func(int, int, int)
	build = func(s, t, p int) {
		if s == t {
			seg.dp[s] = nums[s-1] // nums index start from 0 and dp index start from 1
			return
		}
		m := s + (t-s)>>1
		build(s, m, 2*p)
		build(m+1, t, 2*p+1)
		seg.dp[p] = seg.dp[2*p] + seg.dp[2*p+1]
	}
}

/*
 [l, r] 的区间 都加上 c
 */
func (seg *SegmentTreeB) Add(l, r, s, t, c, p int) {
	if l <= s && t <= r {
		seg.dp[p] += (r - l + 1) * c
		seg.b[p] += c
		return
	}
	m := s + (t-s)>>1
	if seg.b[p] != 0 {
		seg.dp[2*p] += (m - s + 1) * seg.b[p]
		seg.dp[2*p+1] += (t - m) * seg.b[p]
		seg.b[2*p] += seg.b[p]
		seg.b[2*p+1] += seg.b[p]
		seg.b[p] = 0
	}

	if l <= m {
		seg.Add(l, r, s, m, c, 2*p)
	}
	if r > m {
		seg.Add(l, r, m+1, t, c, 2*p+1)
	}
}

func (seg *SegmentTreeB) Getsum(l, r, s, t, p int) int {
	if l <= s && t <= r {
		return seg.dp[p]
	}
	m := s + (t-s)>>1
	if seg.b[p] != 0 {
		seg.dp[2*p] += seg.b[p] * (m - s + 1)
		seg.dp[2*p+1] += seg.b[p] * (t - m)
		seg.b[2*p] += seg.b[p]
		seg.b[2*p+1] += seg.b[p]
		seg.b[p] = 0
	}
	sum := 0
	if l <= m {
		sum += seg.Getsum(l, r, s, m, 2*p)
	}
	if r > m {
		sum += seg.Getsum(l, r, m+1, t, 2*p+1)
	}
	return sum
}
