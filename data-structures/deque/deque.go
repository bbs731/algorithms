package deque

/***
golang 里没有 deque 的数据结构， 可以用 2 个 "头对头"的 slice 来模拟。
主要理解一下， 头对头的含义是什么。
 */

// 在知道数据量的情况下，也可以直接创建个两倍数据量大小的 slice
type Deque struct {
	l, r []any
}

func (q Deque) Empty() bool {
	return len(q.l) == 0 && len(q.r) == 0
}

func (q Deque) Len() int {
	return len(q.l) + len(q.r)
}

// 所有的这些设计就是为了，能在 Front 和 Back 方便的插入，和弹出。
func (q *Deque) PushFront(v any) {
	q.l = append(q.l, v)
}

func (q *Deque) PushBack(v any) {
	q.r = append(q.r, v)
}

// deque 的头再 l 的尾部， 如果 l 是空的， 则头部在 r 的头部
func (q Deque) Front() any {
	if len(q.l) > 0 {
		return q.l[len(q.l)-1]
	}
	return q.r[0]
}

func (q *Deque) PopFront() (v any) {
	if len(q.l) > 0 {
		v, q.l = q.l[len(q.l)-1], q.l[:len(q.l)-1]
	} else {
		v, q.r = q.r[0], q.r[1:]
	}
	return
}

// deque 的尾部， 在 r 的尾部， 如果 r 是空的， 则在 l 的头部。
func (q Deque) Back() any {
	if len(q.r) > 0 {
		return q.r[len(q.r)-1]
	}
	return q.l[0]
}

func (q *Deque) PopBack() (v any) {
	if len(q.r) > 0 {
		v, q.r = q.r[len(q.r)-1], q.r[:len(q.r)-1]
	} else {
		v, q.l = q.l[0], q.l[1:]
	}
	return
}

func (q Deque) Get(i int) any {
	if i < len(q.l) {
		return q.l[len(q.l)-1-i] //这个地方的计算容易错。
	}
	return q.r[i-len(q.l)]
}
