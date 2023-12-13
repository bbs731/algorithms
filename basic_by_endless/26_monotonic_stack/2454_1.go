package monotonic_stack

import "sort"


// 复杂度是 O（n*logn) 会超时
func secondGreaterElement(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	stack1 := make([]int,0)
	stack2 := make([]int, 0)

	for index, x := range nums {
		n2 := sort.Search(len(stack2), func(i int) bool { return nums[stack2[i]] >= x})
		for j :=0; j<n2; j++ {
			ans[stack2[j]] = x
		}
		// pop stack2 and insert x
		stack2 = stack2[n2:]

		n1 := sort.Search(len(stack1), func (i int) bool { return nums[stack1[i]] >= x})
		for j:=0; j<n1; j++ {
			// move candidate from stack1 to stack2
			candidate := nums[stack1[j]]
			pos := sort.Search(len(stack2), func(i int) bool {return nums[stack2[i]] >= candidate})
			if pos == 0 {
				stack2 = append([]int{stack1[j]}, stack2[:]...)
			} else {
				stack2 = append(stack2[:pos], append([]int{stack1[j]}, stack2[pos:]...)...)
			}
		}
		// pop stack1
		stack1 = stack1[n1:]
		stack1 = append([]int{index}, stack1[:]...)
	}

	for j :=0; j < len(stack1); j++ {
		ans[stack1[j]]	= -1
	}
	for j :=0; j<len(stack2); j++ {
		ans[stack2[j]]= -1
	}
	return ans
}


// 你得到了什么教训啊？ 
func secondGreaterElement(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i:=0; i<n; i++ {
		ans[i]= -1
	}

	// stack 单调栈， 队首的元素最大， 队尾的元素最小， 降序的。  //之前考虑的 stack 升序的。
	stack1 := make([]int,0)
	stack2 := make([]int, 0)

	// 处理 stack1， stack2 的时候是有技巧的。 可以简化成 O（1）的操作。
	for index, x := range nums {
		// 先处理 stack2 是对的，
		//n2 := sort.Search(len(stack2), func(i int) bool { return nums[stack2[i]] >= x})
		//for j :=0; j<n2; j++ {
		//	ans[stack2[j]] = x
		//}
		//// pop stack2 and insert x
		//stack2 = stack2[n2:]

		for len(stack2) >0 && x > nums[stack2[len(stack2)-1]] {
			ans[stack2[len(stack2)-1]]	 = x
			// pop stack2
			stack2 = stack2[:len(stack2)-1]
		}
		//n1 := sort.Search(len(stack1), func (i int) bool { return nums[stack1[i]] >= x})
		//for j:=0; j<n1; j++ {
		//	// move candidate from stack1 to stack2
		//	candidate := nums[stack1[j]]
		//	pos := sort.Search(len(stack2), func(i int) bool {return nums[stack2[i]] >= candidate})
		//	if pos == 0 {
		//		stack2 = append([]int{stack1[j]}, stack2[:]...)
		//	} else {
		//		stack2 = append(stack2[:pos], append([]int{stack1[j]}, stack2[pos:]...)...)
		//	}
		//}

		// 处理 stack1 有技巧。
		j := len(stack1) -1
		for j>=0 && x > stack1[j]{
			j--
		}
		// 因为先处理的stack2 所以 stack2 里面的元素，现在都是 >=x 的元素了。 因此，直接可以把stack1 需要去掉的元素，一起转移到 Stack1 就行。
		// 并且保持的是降序。
		stack2 = append(stack2, stack1[j+1:]...)
		stack1 = stack1[:j+1]
		stack1 = append(stack1, index)

		//stack1 = stack1[n1:]
		//stack1 = append([]int{index}, stack1[:]...)
	}


	return ans
}



func secondGreaterElement(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i:=0; i<n; i++ {
		ans[i]= -1
	}

	// stack 单调栈， 队首的元素最大， 队尾的元素最小， 降序的。  //之前考虑的 stack 升序的。
	stack1 := make([]int,0)
	stack2 := make([]int, 0)

	// 处理 stack1， stack2 的时候是有技巧的。 可以简化成 O（1）的操作。
	for index, x := range nums {
		for len(stack2) >0 && x > nums[stack2[len(stack2)-1]] {
			ans[stack2[len(stack2)-1]]	 = x
			// pop stack2
			stack2 = stack2[:len(stack2)-1]
		}
		// 处理 stack1 有技巧。
		j := len(stack1) -1
		for j>=0 && x > nums[stack1[j]]{
			j--
		}
		// 因为先处理的stack2 所以 stack2 里面的元素，现在都是 >=x 的元素了。 因此，直接可以把stack1 需要去掉的元素，一起转移到 Stack1 就行。
		// 并且保持的是降序。
		stack2 = append(stack2, stack1[j+1:]...)
		stack1 = stack1[:j+1]
		stack1 = append(stack1, index)

	}
	return ans
}
