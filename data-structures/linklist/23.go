package linklist

import (
	"container/heap"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type pair struct {
	val int
	i   int
}

type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].val < h[j].val
}

func (h hp) Swap(i, j int) {
	//*h[i], *h[j] = *h[j], *h[i]
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Pop() any {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

func (h *hp) Push(v any) {
	*h = append(*h, v.(pair))
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &hp{}
	heap.Init(h)

	// initialize heap
	for i, l := range lists {
		if l != nil {
			heap.Push(h, pair{l.Val, i})
			//nodes[i] = l
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
