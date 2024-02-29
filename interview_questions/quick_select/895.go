package quick_select

import (
	"container/heap"
	"fmt"
)

/***
设计一个类似堆栈的数据结构，将元素推入堆栈，并从堆栈中弹出出现频率最高的元素。

实现 FreqStack 类:

FreqStack() 构造一个空的堆栈。
void push(int val) 将一个整数 val 压入栈顶。
int pop() 删除并返回堆栈中出现频率最高的元素。
如果出现频率最高的元素不只一个，则移除并返回最接近栈顶的元素。


示例 1：

输入：
["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"],
[[],[5],[7],[5],[7],[4],[5],[],[],[],[]]
输出：[null,null,null,null,null,null,null,5,7,5,4]
解释：
FreqStack = new FreqStack();
freqStack.push (5);//堆栈为 [5]
freqStack.push (7);//堆栈是 [5,7]
freqStack.push (5);//堆栈是 [5,7,5]
freqStack.push (7);//堆栈是 [5,7,5,7]
freqStack.push (4);//堆栈是 [5,7,5,7,4]
freqStack.push (5);//堆栈是 [5,7,5,7,4,5]
freqStack.pop ();//返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,5,7,4]。
freqStack.pop ();//返回 7 ，因为 5 和 7 出现频率最高，但7最接近顶部。堆栈变成 [5,7,5,4]。
freqStack.pop ();//返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,4]。
freqStack.pop ();//返回 4 ，因为 4, 5 和 7 出现频率最高，但 4 是最接近顶部的。堆栈变成 [5,7]。
 */

type tuple struct {
	v    int
	tick []int
	cnts int

	index int
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
