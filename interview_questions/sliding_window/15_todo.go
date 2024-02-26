package sliding_window

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-2; {
		a := nums[i]

		l, r := i+1, n-1
		for l < r {
			b, c := nums[l], nums[r]
			if b+c > -a {
				r--
			} else if b+c < -a {
				l++
			} else {
				ans = append(ans, []int{a, b, c})
				for l++; l < r && l < n-1 && nums[l] == nums[l-1]; l++ {  // 这里比较 n[l], n[l-1]
				}
				for r--; l < r && r < n-1 && nums[r] == nums[r+1]; r-- {  // 这里比较 n[r], n[r+1]
				}
			}
		}
		for i++; i < n-2 && nums[i] == nums[i-1]; i++ {
		}
	}
	return ans
}
