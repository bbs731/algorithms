package interview_questions


func wiggleMaxLength(nums []int) int {
	n := len(nums)

	/***
	  16, - 12, (5, 3, 2), (-5,-5), 11, -8
	  这道题，其实就是在算， 把正数和负数都 group 在一起之后， 总共有多少个 group
	  track nums[i] 和 last 数之间的符号变化就可以。
	  corner case 是处理一下 0 的情况。
	 */
	last := 0
	cnts := 0
	for i :=1; i<n; i++ {
		diff := nums[i] - nums[i-1]
		if diff == 0 {
			continue
		}
		if last == 0 {
			last = diff
			cnts++
			continue
		}
		if last * diff < 0 {
			cnts++
			last = diff
		}
	}
	return cnts + 1 // 这里是 bug 总是忘了加1
}
