package difference

func carPooling(trips [][]int, capacity int) bool {
	n := 0
	diffs := make([]int, 1000+1)

	for _, t := range trips {
		c, start, end := t[0], t[1], t[2]
		diffs[start] += c
		diffs[end] -= c // 为了处理， intervals 相交的问题， 这里不用 diffs[end+1] 反而用 diffs[end] 这是个逻辑技巧。
		n = max(n, end)
	}

	s := 0
	for i := 0; i <= n; i++ {
		s += diffs[i]
		if s > capacity {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
