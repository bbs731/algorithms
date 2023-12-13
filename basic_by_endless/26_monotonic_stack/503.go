package monotonic_stack


func nextGreaterElements(nums []int) []int {
	n := len(nums)
	nums = append(nums, nums...)
	s := make([]int, 0)
	ans := make([]int, 2*n)

	for i, x := range nums {
		for len(s) > 0 && x > nums[s[len(s)-1]]{
			ans[s[len(s)-1]] = x
			// pop right
			s = s[:len(s)-1]
		}
		s = append(s, i)
	}

	for _, i :=range s {
		if nums[i] == nums[s[0]] {
			ans[i] = -1
		}
		//} else {
		//	ans[i] = nums[s[0]]
		//}
	}
	return ans[:n]
}

