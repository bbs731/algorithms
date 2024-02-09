package linklist

import "container/heap"

type pair struct {
	val int
	i   int
}

// 这样写，不是更加的简单吗？
type hp struct{ l []pair }

func (h hp) Len() int {
	return len(h.l)
}

func (h hp) Less(i, j int) bool {
	return h.l[i].val < h.l[j].val
}

// swap 需要的是 hp 不是 *hp 这个让我震惊啊！ heap 果然是死穴啊！
func (h hp) Swap(i, j int) {
	h.l[i], h.l[j] = h.l[j], h.l[i]
}

func (h *hp) Pop() any {
	v := h.l[len(h.l)-1]
	h.l = h.l[:len(h.l)-1]
	return v
}

func (h *hp) Push(v any) {
	h.l = append(h.l, v.(pair))
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &hp{}
	heap.Init(h)

	// initialize heap
	for i, l := range lists {
		if l != nil {
			heap.Push(h, pair{l.Val, i})
		}
	}

	dummy := &ListNode{}
	head := dummy

	for h.Len() > 0 {
		v := heap.Pop(h) // heap 就是你的死穴啊， 这里错了无数次啊。
		i := v.(pair).i
		if lists[i] != nil {
			head.Next = lists[i]
			head = lists[i]
			lists[i] = lists[i].Next
			if lists[i] != nil {
				heap.Push(h, pair{lists[i].Val, i})
			}
		}
	}

	return dummy.Next
}
