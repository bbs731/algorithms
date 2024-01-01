package _600_1900


/*
考虑了 20分钟， 一遍过， 这就是个智力的题目，没有任何技巧！
 */
func maxSubarrays(nums []int) int {
	ans := 0
	res := -1
	for _, num := range nums {
		res = res & num
		if res == 0 {
			ans++
			res = -1
		}
	}
	//if ans == 0 {
	//	return 1
	//}
	//return ans
	return max(ans, 1)
}