package sliding_window

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left := 0
	sum :=0
	ans := n+1

	for right, x :=range  nums {
		sum += x
		for sum >=target {
			ans = min(ans, right -left +1)
			sum -= nums[left]
			left++
		}
	}
	if ans == n+1 {
		return 0
	}
	return ans
}
