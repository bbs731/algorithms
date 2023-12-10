package sliding_window

/*
	经典的 相向双指针的问题
 */
func twoSum(numbers []int, target int) []int {
	n :=len(numbers)
	left := 0
	right := n-1

	for left <= right {
		s := numbers[left]	 + numbers[right]
		if s > target {
			right--
		} else if s < target {
			left++
		}else{
				break
		}
	}
	return []int{left+1, right+1}
}
