package stack

/***
这个解法，也是太优雅了！

 */
type FreqStack struct {
	cnt    map[int]int
	stacks [][]int
}

func Constructor() FreqStack {
	return FreqStack{cnt: map[int]int{}}
}

func (f *FreqStack) Push(val int) {
	c := f.cnt[val]
	if c == len(f.stacks) { // 这个元素的频率已经是目前最多的，现在又出现了一次
		f.stacks = append(f.stacks, []int{val}) // 那么必须创建一个新栈
	} else {
		f.stacks[c] = append(f.stacks[c], val) // 否则就压入对应的栈
	}
	f.cnt[val]++ // 更新频率
}

func (f *FreqStack) Pop() int {
	back := len(f.stacks) - 1
	st := f.stacks[back]
	bk := len(st) - 1
	val := st[bk] // 弹出最右侧栈的栈顶
	if bk == 0 { // 栈为空
		f.stacks = f.stacks[:back] // 删除
	} else {
		f.stacks[back] = st[:bk]
	}
	f.cnt[val]-- // 更新频率
	return val
}
