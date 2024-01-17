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
	rep_cnt  int // 当前这个值（key）重复出现的次数
	//msz      int // 这个是干啥用的？
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
	o.sz = o.rep_cnt + o.lr[0].size() + o.lr[1].size()
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
	// 这是 32-bit 的。 64bit 用的是 <<13, >>7, <<17
	t.random ^= t.random << 13
	t.random ^= t.random >> 17
	t.random ^= t.random << 5
	return t.random
}

func (t *treap) _put(o *tpNode, key int, val int) *tpNode {
	if o == nil {
		return &tpNode{priority: t.fastRand(), sz: 1, key: key, val: val, rep_cnt: 1}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key, val)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else { // d == -1 means same key
		o.rep_cnt++
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
		if o.rep_cnt > 1 {
			o.rep_cnt--
		} else {
			if o.lr[1] == nil {
				return o.lr[0] // 丢弃 root 指向的 o , 让root 指向 o.lr[0]。依赖 golang 的 GC回收 o (原root)
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

// 根据 rank 查询值
// 参考的代码： https://oi-wiki.org/ds/treap/
func (t *treap) _query_val(o *tpNode, rank int) int {
	if o == nil || rank <= 0 || rank > o.size() {
		return -1145 // not found
	}

	s := 0
	if o.lr[0] != nil {
		s = o.lr[0].size()
	}
	if rank == s+1 {
		return o.key
	} else if rank <= s {
		return t._query_val(o.lr[0], rank)
	} else {
		return t._query_val(o.lr[1], rank-s-1)
	}
}

func (t *treap) delete(key int) {
	t.root = t._delete(t.root, key)
}

// 代码缺少了排名树中的, kth() or query_rank 和 query_val() 有例题的时候补上。
func (t *treap) query_val(rank int) int {
	return t._query_val(t.root, rank)
}

func (t *treap) query_rank(key int) int {
	return t._query_rank(t.root, key)
}

func (t *treap) _query_rank(o *tpNode, key int) int {
	s := 0
	if o.lr[0] != nil {
		s = o.lr[0].size()
	}
	if key == o.key {
		return s + 1
	} else if key < o.key {
		if o.lr[0] != nil {
			return t._query_rank(o.lr[0], key)
		} else {
			return 1 // 左子树是空的， 说明要查的 key 比最小的节点还要小
		}
	} else {
		if o.lr[1] != nil {
			return s + 1 + t._query_rank(o.lr[1], key)
		} else {
			return o.size() + 1 // 没有右子树的话，返回整个树 + 1
		}
	}
}

// 查询第一个比 key 小的节点（的key值）
var q_prev_tmp int

func (t *treap) _query_prev(o *tpNode, key int) int {
	if key <= o.key {
		if o.lr[0] != nil {
			return t._query_prev(o.lr[0], key)
		}
	} else {
		// q_prev_tmp 是个全局变量。
		q_prev_tmp = o.key
		if o.lr[1] != nil { // 如果有右子树，就一直向右边走。
			t._query_prev(o.lr[1], key)
		}
		return q_prev_tmp
	}

	return -1145
}

var q_next_tmp int

// 查询第一个比 key 大的节点
func (t *treap) _query_next(o *tpNode, key int) int {
	if key >= o.key {
		if o.lr[1] != nil {
			return t._query_next(o.lr[1], key)
		}
	} else {
		q_next_tmp = o.key
		if o.lr[0] != nil {
			t._query_next(o.lr[0], key)
		}
		return q_next_tmp
	}
	return -1145
}
