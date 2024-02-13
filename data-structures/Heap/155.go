package Heap

type MinStack struct {
	st []int
	m  []int // m 是一个单调栈，栈顶元素最小。
}

func Constructor() MinStack {
	st := make([]int, 0)
	m := make([]int, 0)

	return MinStack{
		st,
		m,
	}
}

func (this *MinStack) Push(val int) {
	this.st = append(this.st, val)
	if len(this.m) > 0 && val > this.m[len(this.m)-1] {
		// nothing to do
		return
	}
	// otherwise push the smallest element on top
	this.m = append(this.m, val)
}

func (this *MinStack) Pop() {
	v := this.st[len(this.st)-1]
	this.st = this.st[:len(this.st)-1]
	if len(this.m) > 0 && v == this.m[len(this.m)-1] {
		// pop m
		this.m = this.m[:len(this.m)-1]
	}
}

func (this *MinStack) Top() int {
	return this.st[len(this.st)-1]
}

func (this *MinStack) GetMin() int {
	return this.m[len(this.m)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
