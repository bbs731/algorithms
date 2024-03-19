package montonic_stack

func maximumScore(nums []int, k int) int {
	n := len(nums)
	st := []int{}
	left := make([]int, n)
	right := make([]int, n)
	for i := range left {
		left[i] = -1
		right[i] = n
	}

	ans := 0
	for i, num := range nums {
		for len(st) > 0 && num < nums[st[len(st)-1]] {
			top := st[len(st)-1]
			right[top] = i
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1]
		}
		st = append(st, i)
	}

	for i, num := range nums {
		if left[i]+1 <= k && right[i]-1 >= k {
			ans = max(ans, (right[i]-1-left[i])*num)
		}
	}
	return ans
}
