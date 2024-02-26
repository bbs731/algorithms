package sliding_window

func twoSum(numbers []int, target int) (ans []int) {
	l, r := 0, len(numbers)-1

	for l < r {
		a, b := numbers[l], numbers[r]
		if a+b > target {
			r--
		} else if a+b < target {
			l++
		} else {
			ans = []int{l + 1, r + 1}
			return
		}
	}
	return []int{}
}
