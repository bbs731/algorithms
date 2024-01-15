package weekly

import "sort"




func prefix_function(pattern string) []int {
	n := len(pattern)
	pi := make([]int, n)

	j := 0 // j 记录的是 pi[i-1], 初始化为 pi[0]  即为 0
	for i := 1; i < n; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
		}

		if pattern[i] == pattern[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}

func kmp(text, pattern string) []int {

	pi := prefix_function(pattern)
	pos := make([]int, 0)

	j := 0
	for i := 0; i < len(text); i++ {
		for j > 0 && pattern[j] != text[i] {
			j = pi[j-1]
		}
		if pattern[j] == text[i] {
			j++
		}
		if j == len(pattern) {
			pos = append(pos, i-len(pattern)+1)
			j = pi[j-1]
		}
	}
	return pos
}

func abs (a, b int)int {
	if a > b {
		return a - b
	}
	return b - a
}

func beautifulIndices(s string, a string, b string, k int) []int {
	apos := kmp(s, a)
	bpos := kmp(s, b)

	ans := []int{}
	root := newStRoot(0, 1e9)

	for _, bp := range bpos {
		root.update(bp, 1)
	}

	for _, ap := range apos {
		// 哈哈，真是来了一把牛刀！
		cnt := root.query(max(0,ap-k), min(ap+k, len(b)-1))
		if cnt >=1 {
			ans = append(ans, ap)
		}
	}
	return ans
}


const stNodeDefaultVal = -1 // 如果求最大值并且有负数，改成 math.MinInt

type stNode struct {
	lo, ro *stNode
	l, r   int
	val    int
}

var emptyStNode = &stNode{val: stNodeDefaultVal}

func init() {
	emptyStNode.lo = emptyStNode
	emptyStNode.ro = emptyStNode
}

// 0 1e9
// -2e9 2e9
func newStRoot(l, r int) *stNode {
	return &stNode{lo: emptyStNode, ro: emptyStNode, l: l, r: r, val: stNodeDefaultVal}
}

func (stNode) mergeInfo(a, b int) int {
	return max(a, b)
}

func (o *stNode) maintain() {
	o.val = o.mergeInfo(o.lo.val, o.ro.val)
}

func (o *stNode) update(i, val int) {
	if o.l == o.r {
		o.val = o.mergeInfo(o.val, val)
		return
	}
	m := (o.l + o.r) >> 1
	if i <= m {
		if o.lo == emptyStNode {
			o.lo = &stNode{lo: emptyStNode, ro: emptyStNode, l: o.l, r: m, val: stNodeDefaultVal}
		}
		o.lo.update(i, val)
	} else {
		if o.ro == emptyStNode {
			o.ro = &stNode{lo: emptyStNode, ro: emptyStNode, l: m + 1, r: o.r, val: stNodeDefaultVal}
		}
		o.ro.update(i, val)
	}
	o.maintain()
}

func (o *stNode) query(l, r int) int {
	if o == emptyStNode || l > o.r || r < o.l {
		return stNodeDefaultVal
	}
	if l <= o.l && o.r <= r {
		return o.val
	}
	return o.mergeInfo(o.lo.query(l, r), o.ro.query(l, r))
}

// end of 开点线段树