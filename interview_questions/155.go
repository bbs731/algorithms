package interview_questions

type pair struct {
	val      int
	minvalue int
}
type MinStack struct {
	st []pair
	//m  []int // m 是一个单调栈，栈顶元素最小。
}

func Constructor() MinStack {
	st := make([]pair, 0)
	//m := make([]int, 0)

	return MinStack{
		st,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func (this *MinStack) Push(val int) {
	if len(this.st) == 0 {
		this.st = append(this.st, pair{val, val})
		return
	}
	// otherwise push the smallest element on top
	this.st = append(this.st, pair{val, min(val, this.st[len(this.st)-1].minvalue)})
}

func (this *MinStack) Pop() {
	this.st = this.st[:len(this.st)-1]
}

func (this *MinStack) Top() int {
	return this.st[len(this.st)-1].val
}

func (this *MinStack) GetMin() int {
	return this.st[len(this.st)-1].minvalue
}

