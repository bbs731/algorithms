package montonic_stack

func largestRectangleArea(heights []int) int {
	n := len(heights)
	st := []int{}
	left := make([]int, n)
	right := make([]int, n)
	for i := range left {
		// 胆子够大的是吧。  这里初始化， 或者 让  st := []int{-1} 这样也是可以的。
		left[i] = -1
		right[i] = n
	}

	for i := 0; i < n; i++ {
		for len(st) > 0 && heights[i] <= heights[st[len(st)-1]] {
			right[st[len(st)-1]] = i
			// pop st
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1]
		}
		// push into stack
		st = append(st, i)
	}

	ans := 0
	for i := 0; i < n; i++ {
		l, r := left[i], right[i]
		ans = max(ans, heights[i])
		ans = max(ans, (r-l-1)*heights[i])
	}
	return ans
}
