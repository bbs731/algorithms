package montonic_stack

/***

以2104的题目，作为 monotonic stack template code 加以说明。
 */
func solve(nums []int) (ans int64) {
	n := len(nums)
	left := make([]int, n)  // left[i] 为左侧严格大于 num[i] 的最近元素位置（不存在时为 -1）
	right := make([]int, n) // right[i] 为右侧大于等于 num[i] 的最近元素位置（不存在时为 n）
	for i := range right {
		right[i] = n // put dummy boundary
	}
	st := []int{-1} // put dummy boundary
	for i, v := range nums {

		// 单调栈的性质主要给了我们两个信息：
		// 1. 告诉了我们， 对于 当前的数 v (第 ith element), 往左看，第一个比当前 v 大 (或者小， 对于本题是大）的数的位置。
		// 2. 随便的在 pop stack element 的时候，我们知道， ith element 是这些在stack上即将被pop 元素的，右边的第一个比他们大的元素就是v . 这就是为什么，单调栈可以用来求，下一个更大元素的题目的原因。
		for len(st) > 1 && nums[st[len(st)-1]] <= v {
			right[st[len(st)-1]] = i // 2. 的信息。
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1] // 1. 的信息。
		st = append(st, i)
	}

	// 统计答案的工作， 每道题，会略有不同。
	for i, v := range nums {
		ans += (int64(i-left[i])*int64(right[i]-i) - 1) * int64(v)
	}
	return
}
