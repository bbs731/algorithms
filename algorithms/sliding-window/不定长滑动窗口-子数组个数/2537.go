package sliding_window


// 想法太牛逼了， 一次就过了。
// 估计过几个月，有开始懵逼，会不会?
func countGood(nums []int, k int) int64 {
	n := len(nums)
	counter := make(map[int]int)
	left :=0
	ans := 0
	pairs := 0

	for i, v := range nums {
		counter[v]++
		pairs += counter[v]-1

		for pairs >=k {
			ans += n-i  // 你他妈的这里算是学通了是吧！
			// 这里计数， 已做端点作为参考点， 从 [i, n-1] 总共有  n-i 个子集是满足要求的，加入到 ans
			// 下面的 left++
			// try to move the left
			l := nums[left]

			counter[l]--
			pairs -= counter[l]
			left++
		}
	}

	return int64(ans)
}
