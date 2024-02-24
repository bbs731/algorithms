package hash


/***
这道题， 有意思， 感觉，是442 这道题的反命题，哈哈！
 */
func findDisappearedNumbers(nums []int) (ans []int) {

	for i :=range nums {
		for nums[nums[i]-1] != nums[i]{
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i, num :=range nums {
		if num-1 != i {
			ans = append(ans, i+1)  // 这里用 i+1 就是那个 missing 的 digit
		}
	}
	return
}
