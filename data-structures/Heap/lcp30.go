package Heap

import (
	"container/heap"
	"sort"
)

func magicTower(nums []int) int {
	n := len(nums)

	h := &hp{}
	heap.Init(h)

	sum := 0
	cnts := 0
	for i := 0; i < n-1; i++ {
		num := nums[i]
		// heap 的正确使用姿势
		heap.Push(h, num)
		sum += num

		for sum < 0 {
			// heap 的正确使用姿势
			v := heap.Pop(h)
			cnts++
			sum -= v.(int)       // v is negative
			nums[n-1] += v.(int) // negative accumulated to nums[n-1]
		}
	}
	sum += nums[n-1]
	if sum < 0 {
		return -1
	}
	return cnts
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
