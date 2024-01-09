package Heap

import (
	"container/heap"
	"sort"
)

// 还有个简单的写法， 利用 sort.IntSlice  这样就不用写 Less, Len, 和 Swap 了。
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // > 为最大堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// 注意，这里需要使用 pointer type *IntHeap
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func main() {
	h := &IntHeap{}
	heap.Init(h)
	heap.Push(h, 3)
	item := heap.Pop(h).(int)
	_ := item
}

// 利用 sort.IntSlice 写 int heap
type hp struct{ sort.IntSlice }

// 因为 sort.IntSlice， 可以省去 less, len, swap 的代码
func (h *hp) Push(v any) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() any {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

// 还有带信息的 heap, 遇到问题再添加
