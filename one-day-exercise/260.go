package one_day_exercise

func singleNumber(nums []int) []int {
	total := 0
	for _, x := range nums {
		total ^= x
	}
	// find the last bit
	//lowbit := total & (^total + 1) // x &(^x +1) 取最后一个是1的bit
	lowbit := total & (-total) // -total = ^total + 1  你注意到这一点了吗？

	one, two := 0, 0
	for _, x := range nums {
		if x&lowbit > 0 {
			one ^= x
		} else {
			two ^= x
		}
	}
	return []int{one, two}

}
