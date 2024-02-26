package sliding_window

func countPairs(nums []int, target int) int {
	ans := 0
	n := len(nums)
	sort.Ints(nums)

	l, r := 0, n-1
	for l < r {
		if nums[l]+nums[r] < target {
			ans += r - l
			l++
		} else {
			r--
		}
	}
	return ans
}
