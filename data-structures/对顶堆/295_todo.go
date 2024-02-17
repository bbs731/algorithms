package head_to_head_heap

import "container/heap"

/***
对顶堆栈的，模版题目。

进阶题目，就是， 对顶堆 + sliding window  这时候需要 lazy heap 更好的替代品，应该是 red-black tree
 */
type MedianFinder struct {
	small hp   // 最大堆
	big hp     // 最小堆， 放入负数
}

func Constructor() MedianFinder {
	small := make(hp, 0)
	big := make(hp, 0)
	heap.Init(&small)
	heap.Init(&big)
	return MedianFinder{
		small,
		big,
	}
}

func (this *MedianFinder) AddNum(num int)  {
	sl , bl := this.small.Len(), this.big.Len()
	if sl > bl {  // sl = bl + 1
		heap.Push(&this.small, num)
		v := heap.Pop(&this.small).(int)
		// insert into big
		heap.Push(&this.big, -v)
	}else { // sl == bl
		heap.Push(&this.big, -num)
		v := heap.Pop(&this.big).(int)
		// insert into samll
		heap.Push(&this.small, -v)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	sl, bl := this.small.Len(), this.big.Len()

	if sl > bl {
		return float64(this.small[0])
	}
	// sl == bl
	return float64(this.small[0] - this.big[0])/2
}


type hp []int
func (h hp) Len() int {
	return len(h)
}

func (h hp) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}

func (h hp) Less (i, j int) bool {
	return h[i] > h[j]    //最大堆
}

func (h *hp)Push(v any){
	*h = append(*h, v.(int))
}

func (h *hp) Pop() any {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */