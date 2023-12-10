package monotonic_stack


func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stack := make([]int, 0)  //monotoinic stack 里面放的是 index
	ans := make([]int, n)

	for i, x := range temperatures {
		for len(stack)> 0 && x > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			ans[top] = i - top
			// pop top
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}


func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stack := make([]int, 0)  //monotoinic stack 里面放的是 index
	ans := make([]int, n)

	for i, x := range temperatures {
		for len(stack)> 0 {
			top := stack[len(stack)-1]
			if temperatures[top] >= x {
				break
			} else {
				ans[top] = i - top
				// pop top
				stack = stack[:len(stack)-1]
			}
		}
		// push x
		stack = append(stack, i)
	}
	return ans
}
