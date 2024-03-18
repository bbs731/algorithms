package stack


type MyQueue struct {
	st1 []int
	st2 []int
}


func Constructor() MyQueue {
	return MyQueue {
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *MyQueue) Push(x int)  {
	this.st1 = append(this.st1, x)
}


func (this *MyQueue) Pop() int {
	if len(this.st2) == 0 {
		for len(this.st1) > 0 {
			this.st2 = append(this.st2, this.st1[len(this.st1)-1])
			this.st1 = this.st1[:len(this.st1)-1]
		}
	}
	v := this.st2[len(this.st2)-1]
	this.st2 = this.st2[:len(this.st2)-1]
	return v
}


func (this *MyQueue) Peek() int {
	if len(this.st2) == 0 {
		for len(this.st1) > 0 {
			this.st2 = append(this.st2, this.st1[len(this.st1)-1])
			this.st1 = this.st1[:len(this.st1)-1]
		}
	}
	return this.st2[len(this.st2)-1]
}


func (this *MyQueue) Empty() bool {
	return len(this.st2)== 0 && len(this.st1) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */