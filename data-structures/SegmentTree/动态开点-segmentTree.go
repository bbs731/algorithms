package SegmentTree

//开点线段树，通常会运用在 值域上。 譬如 （-2e9,  2e9)

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

// end of 开点线段树(单点修改）

// 动态开点线段树·其二·延迟标记（区间修改）
var lazyRoot = &lazyNode{l: 1, r: 1e9, sum: stNodeDefaultVal}

const stNodeDefaultTodoVal = 0

type lazyNode struct {
	lo, ro *lazyNode
	l, r   int
	sum    int
	todo   int
}

func (o *lazyNode) get() int {
	if o != nil {
		return o.sum
	}
	return stNodeDefaultVal
}

func (lazyNode) op(a, b int) int {
	return a + b // max(a, b)
}

func (o *lazyNode) maintain() {
	o.sum = o.op(o.lo.get(), o.ro.get())
}

// 没试过， 这个build，是怎么用的
func (o *lazyNode) build(a []int, l, r int) {
	o.l, o.r = l, r
	o.todo = stNodeDefaultTodoVal
	if l == r {
		o.sum = a[l-1]
		return
	}
	m := (l + r) >> 1
	o.lo = &lazyNode{}
	o.lo.build(a, l, m)
	o.ro = &lazyNode{}
	o.ro.build(a, m+1, r)
	o.maintain()
}

func (o *lazyNode) do(add int) {
	o.todo += add                  // % mod
	o.sum += (o.r - o.l + 1) * add // % mod
}

func (o *lazyNode) spread() {
	m := (o.l + o.r) >> 1
	if o.lo == nil {
		o.lo = &lazyNode{l: o.l, r: m, sum: stNodeDefaultVal}
	}
	if o.ro == nil {
		o.ro = &lazyNode{l: m + 1, r: o.r, sum: stNodeDefaultVal}
	}
	if todo := o.todo; todo != stNodeDefaultTodoVal {
		o.lo.do(todo)
		o.ro.do(todo)
		o.todo = stNodeDefaultTodoVal
	}
}

func (o *lazyNode) update(l, r int, add int) {
	if l <= o.l && o.r <= r {
		o.do(add)
		return
	}
	o.spread()
	m := (o.l + o.r) >> 1
	if l <= m {
		o.lo.update(l, r, add)
	}
	if m < r {
		o.ro.update(l, r, add)
	}
	o.maintain()
}

func (o *lazyNode) query(l, r int) int {
	if o == nil || l > o.r || r < o.l {
		return stNodeDefaultVal
	}
	if l <= o.l && o.r <= r {
		return o.sum
	}
	o.spread()
	return o.op(o.lo.query(l, r), o.ro.query(l, r))
}

const stNodeDefaultVal = 0 // 如果求最大值并且有负数，改成 math.MinInt

// 结束 动态开点线段树·其二·延迟标记（区间修改）
