package Treap

import "time"

/*
灵神的实现
https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/treap.go
 */
type tpNode struct {
	lr       [2]*tpNode // 0:左二子，  1：右儿子
	priority uint       // max heap, root node has hte highest priority value
	sz       int        // 实现排名树
	key      int
	val      int
	msz      int // 这个是干啥用的？
}

func (o *tpNode) cmp(a int) int {
	b := o.key

	if a == b {
		return -1
	}
	if a < b {
		return 0 //左二子
	}
	return 1 //右儿子
}

func (o *tpNode) size() int {
	if o == nil {
		return 0
	}
	return o.sz
}

// 也有叫 update() 的。
func (o *tpNode) maintain() {
	o.sz = 1 + o.lr[0].size() + o.lr[1].size()
}

// 这里是个标准的模板代码
// d=0, 左旋， 返回 o 的右儿子
// d=1, 右旋， 返回 o 的左二子
func (o *tpNode) rotate(d int) *tpNode {
	x := o.lr[d^1] // d^1  =  1-d
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o

	o.maintain()
	x.maintain()
	return x
}

type treap struct {
	random uint
	root   *tpNode
}

func newTreap() *treap {
	return &treap{random: uint(time.Now().UnixNano())/2 + 1}
}

// https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf
// https://en.wikipedia.org/wiki/Xorshift
func (t *treap) fastRand() uint {
	// 这是什么意思？ 查看上面的链接
	t.random ^= t.random << 13
	t.random ^= t.random >> 17
	t.random ^= t.random << 5
	return t.random
}

func (t *treap) _put(o *tpNode, key int, val int) *tpNode {
	if o == nil {
		return &tpNode{priority: t.fastRand(), sz: 1, key: key, val: val}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key, val)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else { // d == -1 means same key
		o.val += val
	}
	o.maintain()
	return o
}

func (t *treap) put(key int, val int) {
	t.root = t._put(t.root, key, val)
}

func (t *treap) _delete(o *tpNode, key int) *tpNode {
	if o == nil {
		return nil
	}

	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._delete(o.lr[d], key)
	} else {
		if o.val > 1 {
			o.val--
		} else {
			if o.lr[1] == nil {
				return o.lr[0]
			}
			if o.lr[0] == nil {
				return o.lr[1]
			}
			d = 0
			if o.lr[0].priority > o.lr[1].priority {
				d = 1
			}
			o = o.rotate(d)
			o.lr[d] = t._delete(o.lr[d], key)
		}
	}
	o.maintain()
	return o
}

func (t *treap) delete(key int) {
	t.root = t._delete(t.root, key)
}

// 代码缺少了排名树中的, kth() 和 find() 有例题的时候补上。
