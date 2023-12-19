package SegmentTree
//李熠东 《算法竞赛进阶指南》

var a []int
const size = 100
type SegmentTree struct {
	l, r int
	dat int  // 最大值
}
var t [size*4]SegmentTree

// p 是 node 的节点
func build(p, l, r int) {
	t[p].l, t[p].r = l, r
	if l == r {
		t[p].dat = a[l]
		return
	}
	mid := l + (r-l)>>1
	build(2*p, l, mid)
	build(2*p+1, mid+1, r)
	t[p].dat = max(t[p*2].dat, t[p*2+1].dat)
}

func change(p, x, v int) {
	if t[p].l == t[p].r {
		t[p].dat = v
		return
	}
	mid := (t[p].l + t[p].r)/2
	if x <= mid {
		change(2*p, x, v )
	} else {
		change (2*p+1, x,v)
	}
	t[p].dat = max(t[p*2].dat, t[p*2+1].dat)
}

func ask(p, l, r int ) int {
	inf := int(1e9)
	if l <=t[p].l && r >=t[p].r{
		return t[p].dat
	}
	mid := (t[p].l + t[p].r)/2
	val := -inf
	if l<=mid {
		val = max(val, ask(p*2, l, r))
	}
	if r >mid {
		val = max(val, ask(p*2+1, l, r))
	}
	return val
}


