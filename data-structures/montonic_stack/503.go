package montonic_stack


func nextGreaterElements(nums []int) []int {
	n := len(nums)
	s := make([]int, 0)
	ans := make([]int, n)
	for i:= range ans {
		ans[i] = -1
	}
	for i:=0; i< 2*n; i++ {
		x := nums[i%n]  // 哈哈， 有技巧  i%n
		for len(s) > 0 && x > nums[s[len(s)-1]]{
			ans[s[len(s)-1]] = x
			s = s[:len(s)-1]
		}
		s = append(s, i%n)
	}
	return ans
}


/***
不用真的那么实在， 开 2*n 的数组， 并且 nums = append(nums, nums..)
直接 loop [0, 2*n-1] 用 i%n 就可以。
 */
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	nums = append(nums, nums...)
	s := make([]int, 0)
	ans := make([]int, 2*n)
	for i:= range ans {
		ans[i] = -1
	}

	for i, x := range nums {
		for len(s) > 0 && x > nums[s[len(s)-1]]{
			ans[s[len(s)-1]] = x
			s = s[:len(s)-1]
		}
		s = append(s, i)
	}
	return ans[:n]
}
