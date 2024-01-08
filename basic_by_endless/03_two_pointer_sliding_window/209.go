package _3_two_pointer_sliding_window

func minSubArrayLen(target int, nums []int) int {
	left := 0
	n := len(nums)
	ans := n + 1
	sum := 0
	for right, x := range nums {
		sum += x
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if ans == n+1 {
		return 0
	}
	return ans
}
