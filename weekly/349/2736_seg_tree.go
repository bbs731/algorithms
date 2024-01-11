package weekly

import (
	"cmp"
	"slices"
)

// start of 开点线段树(支持单点修改）

// 开点线段树的代码来自：
//https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/segment_tree.go#L455
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

func maximumSumQueries(nums1, nums2 []int, queries [][]int) []int {
	type pair struct{ x, y int }
	a := make([]pair, len(nums1))
	for i, x := range nums1 {
		a[i] = pair{x, nums2[i]}
	}
	slices.SortFunc(a, func(a, b pair) int { return cmp.Compare(b.x, a.x) })
	qid := make([]int, len(queries))
	for i := range qid {
		qid[i] = i
	}
	slices.SortFunc(qid, func(i, j int) int { return cmp.Compare(queries[j][0], queries[i][0]) })

	ans := make([]int, len(queries))
	j := 0
	root := newStRoot(1, 1e9)
	for _, i := range qid {
		x, y := queries[i][0], queries[i][1]

		// 目的是找寻满足  a[j].x >=x  一段数组中， a[j].x + a[j]y 的最大值， 这是 range query 的问题。 可以用 BIT, segment tree， STable 解决。
		// 同时涉及到了，需要单点修改，所以淘汰 ST。
		// 这个代码中， 用了支持单点修改的 开点SegmentTree 的实现
		// 因为是 开点SegmentTree 所以， 可以省去，值域的离散化， 看BIT版本的实现。
		for ; j < len(a) && a[j].x >= x; j++ { // 下面只需关心 a[j].y
			root.update(a[j].y, a[j].x+a[j].y)
		}
		ans[i] = root.query(y, 1e9)
	}
	return ans
}
