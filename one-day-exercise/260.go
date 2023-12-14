package one_day_exercise

func singleNumber(nums []int) []int {
	total := 0
	for _, x := range nums {
		total ^= x
	}
	// find the last bit
	lastbit := total & (^total + 1) // x &(^x +1) 取最后一个是1的bit

	one, two := 0, 0
	for _, x := range nums {
		if x&lastbit > 0 {
			one ^= x
		} else {
			two ^= x
		}
	}
	return []int{one, two}

}
