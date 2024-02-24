package hash


/****
还有这个技巧吗？ 叫原地hash

就是说 1...n  这些数字。  1 的位置应该在 0，  n  的位置应该在 n-1

我操， 这也太神奇了！
 */



func findDuplicates(nums []int) []int {
	ans := []int{}

	for i :=range nums{
		for nums[nums[i]-1] != nums[i]{
			//swap
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i, num :=range nums {
		if num-1!= i{
			ans = append(ans, num)
		}
	}
	return ans
}