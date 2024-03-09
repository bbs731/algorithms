package heap

import (
	"container/heap"
	"sort"
)

/***

抄答案，抄答案！ 哎！


这个是一维数组的， 还有二维数组的情况。

 */

type pair struct {
	val int
	i   int
}

func kSum(nums []int, k int) int64 {
	s := 0
	n := len(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			s += nums[i]
		} else {
			nums[i] = -nums[i]
		}
	}
	sort.Ints(nums)
	h := &hp{}
	heap.Init(h)
	h.Push(pair{0, 0})
	for ; k > 1; k-- {
		top := heap.Pop(h).(pair)
		if top.i < n {
			if top.i > 0 {
				heap.Push(h, pair{top.val - nums[top.i-1] + nums[top.i], top.i + 1})
			}
			heap.Push(h, pair{top.val + nums[top.i], top.i + 1})
		}
	}

	return int64(s - (*h)[0].val)
}

type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hp) Less(i, j int) bool {
	return h[i].val < h[j].val
}

func (h *hp) Push(v any) {
	*h = append(*h, v.(pair))
}

func (h *hp) Pop() (v any) {
	a := *h
	*h, v = a[:len(a)-1], a[len(a)-1]
	return
}
