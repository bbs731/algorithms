package dp

/**
看这代码， 感觉功力 长进了， 但是有个问题， 就是， 你一次写不对啊！
 */
func maxProduct(nums []int) int {
	fn, fp, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		fp, fn = max(num*fp, num*fn, num), min(num*fp, num*fn, num)
		if fp < 0 {
			fp = 0
		}

		if fn > 0 {
			fn = 0
		}
		ans = max(ans, fp)
	}
	return ans
}
