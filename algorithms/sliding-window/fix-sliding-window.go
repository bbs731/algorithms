package sliding_window

func fixed_siding_window_template_code(k int, nums []int) int {

	s := 0
	// step1: 先处理  k-1 个 element
	for _, x := range nums[:k] {
		s += x
	}
	ans := s

	for i, num := range nums[k-1:] {
		// kth element 进入窗口
		s += nums[num]
		// 处理 size of k windows elements
		ans = max(ans, s)

		// i-k+1 element 离开窗口
		s -= nums[i-k+1]
	}
	return ans
}
