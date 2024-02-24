package interview_questions


/***
接雨水啊，接雨水

单调栈的解法，还是不熟悉啊！
 */

func trap(height []int) int {
	n :=len(height)
	st :=[]int{}
	left := make([]int, n)
	right :=make([]int, n)
	for i:=0; i<n; i++ {
		left[i]= -1
		right[i]= n
	}

	ans := 0
	for i:=0; i<n; i++ {
		for len(st) >0 && height[i]	>= height[st[len(st)-1]] {
			top := st[len(st)-1]
			right[top] = i
			if left[top] != -1 {
				l, r := left[top], right[top]
				ans += (min(height[l], height[r]) - height[top]) * (r-l-1)
			}
			// pop stack
			st = st[:len(st)-1]
		}
		if len(st)> 0 {
			left[i] = st[len(st)-1]
		}
		st = append(st, i)
	}

	return ans
}


