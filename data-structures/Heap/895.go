package Heap

import "container/heap"

/****
用 heap 来求解，也是可以的。 但是时间复杂度是 O（lgn) 的， 不像 灵神的题解， 简单，而且，时间复杂度是 O(1) 的。
 */
type tuple struct {
	v    int
	tick []int
	cnts int

	index int // 这个解法，在工程的项目上看到过， 需要保存 index 的数值， 以便  heap.Fix(&h, index) 当，当前的 tuple 的 cnts 值需要更新的时候。
}

type FreqStack struct {
	h    hp
	m    map[int]*tuple
	tick int
}

func Constructor() FreqStack {
	h := make(hp, 0)
	heap.Init(&h)
	return FreqStack{h, make(map[int]*tuple), 0}
}

func (this *FreqStack) Push(val int) {
	this.tick++
	if existing, ok := this.m[val]; ok {
		existing.cnts++
		existing.tick = append(existing.tick, this.tick)
		heap.Fix(&this.h, existing.index)
	} else {
		// new entry
		entry := &tuple{
			val,
			[]int{this.tick},
			1,
			-1,
		}
		this.m[val] = entry
		heap.Push(&this.h, entry)
	}
}

func (this *FreqStack) Pop() int {

	v := this.h[0]
	ans := v.v
	fmt.Println(ans)
	if this.m[v.v].cnts == 1 {
		// need to pop
		heap.Pop(&this.h)
		// remove from map
		delete(this.m, v.v)
	} else {
		// decrement cnts
		v.cnts--
		v.tick = v.tick[:len(v.tick)-1]
		heap.Fix(&this.h, v.index)
	}
	return ans
}

type hp []*tuple

func (h hp) Less(i, j int) bool {
	// 最大堆
	return h[i].cnts > h[j].cnts || (h[i].cnts == h[j].cnts && h[i].tick[h[i].cnts-1] > h[j].tick[h[i].cnts-1])
}

func (h hp) Len() int {
	return len(h)
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]

	h[j].index = j
	h[i].index = i
}

func (h *hp) Push(v any) {
	n := len(*h)
	v.(*tuple).index = n
	*h = append(*h, v.(*tuple))
}

func (h *hp) Pop() any {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	v.index = -1
	return v
}
