package mono_queue

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0)
	q := []int{} //存储的是 nums's index
	for i, x := range nums {
		// 入队
		for len(q)>0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1]
		}
		q = append(q, i)

		// 出队
		if i - q[0] +1 > k {
			q = q[1:]
		}

		// 记录结果
		if i >= k -1 {
			ans = append(ans, nums[q[0]])
		}
	}
	return ans
}
