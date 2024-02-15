package interview_questions

func rotate(nums []int, k int)  {
	n := len(nums)
	if n == 0 || k%n == 0 {
		return
	}
	k = k%n

	tmp := make([]int, n)
	for i:=0; i<n; i++ {
		tmp[(i+k)%n] = nums[i]
	}
	copy(nums, tmp)
}

// 牛逼
func rotate(nums []int, k int)  {
	k = k % len(nums)
	nums1 := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	copy(nums, nums1)
}