package SegmentTree

type SegmentTree struct {
	n  int
	dp []int
}

func (seg *SegmentTree) Init(nums []int) {
	n := len(nums)
	seg.n = n
	seg.dp = make([]int, 4*n)

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

func (seg *SegmentTree) Add(l, r, s, t, c, p int) {
	if s == t { // 注意， 因为没有 b[] 惰性的 marker, 这里，只能递推到  s == t 才能更新
		seg.dp[p] += c
		return
	}
	m := s + (t-s)>>1
	if l <= m {
		seg.Add(l, r, s, m, c, 2*p)
	}
	if r > m {
		seg.Add(l, r, m+1, t, c, 2*p+1)
	}
	seg.dp[p] = seg.dp[2*p] + seg.dp[2*p+1]
}

/*
range sum query [l, r]
 */
func (seg *SegmentTree) Getsum(l, r, s, t, p int) int {
	if s == t {
		return seg.dp[p]
	}
	sum := 0
	m := s + (t-s)>>1
	if l <= m {
		sum += seg.Getsum(l, r, s, m, 2*p)
	}
	if r > m {
		sum += seg.Getsum(l, r, m+1, t, 2*p+1)
	}
	return sum
}

/*
range max query [l, r]
 */
func (seg *SegmentTree) GetMax(l, r, s, t, p int) int {
	if s == t {
		return seg.dp[p]
	}
	value := 0
	m := s + (t-s)>>1
	if l <= m {
		value = max(value, seg.GetMax(l, r, s, m, 2*p))
	}
	if r > m {
		value = max(value, seg.GetMax(l, r, m+1, t, 2*p+1))
	}
	return value
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
