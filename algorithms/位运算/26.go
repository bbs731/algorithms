package bits_operation

/**
这个题目和 bits operation 没有关系，单纯的处于理解 2411 模板题目的需要
锻炼一下原地修改。
 */
func removeDuplicates(nums []int) int {
	n := len(nums)
	k := 0
	for i := 1; i < n; i++ {
		if nums[i-1] != nums[i] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}
