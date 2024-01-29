package monotonic_stack

/***
都是基于 84题的应用。
 */
func u84 (heights []int) int {
	n :=len(heights)
	left := make([]int, n)
	right := make([]int, n)
	for i := range right {
		right[i]= n
	}

	st := []int{-1}
	for i, v := range heights {
		for len(st)	 > 1 && v < heights[st[len(st)-1]] {
			right[st[len(st)-1]] = i
			// pop stack
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}

	ans := 0
	for i, v := range heights {
		ans = max(ans, (right[i] - left[i]-1)*v)
	}
	return ans
}

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])

	heights := make([]int, n)
	ans := 0


	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			if matrix[i][j] == '0' {
				heights[j]	= 0
			} else {
				heights[j] += 1
			}
		}
		ans = max(ans, u84(heights))
	}
	return ans
}

