package montonic_queue

/***
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。



示例 1：

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
示例 2：

输入：nums = [1], k = 1
输出：[1]


提示：

1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length

 */

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	q := []int{}

	for i, num := range nums {
		for len(q) > 0 && nums[q[len(q)-1]] <= num {
			q = q[:len(q)-1]
		}
		q = append(q, i)

		// check front, pop out of window index
		for len(q) > 0 && q[0] < i-k+1 { // i-k+1 这个边界，太容易错了！
			q = q[1:]
		}
		if i-k+1 >= 0 {
			ans[i-k+1] = nums[q[0]]
		}
	}
	return ans
}

/***
这个版本， 更自然吗？
 */
func maxSlidingWindow(nums []int, k int) []int {
	st := []int{}
	n := len(nums)
	l := make([]int, 0)
	left := 0
	for i := 0; i < n; i++ {
		for len(st) > 0 && nums[i] > nums[st[len(st)-1]] {
			// pop stack
			st = st[:len(st)-1]
		}
		st = append(st, i)

		// need to move left
		if i-left+1 > k {
			if st[0] == left {
				// need to pop stack as well
				st = st[1:]
			}
			left++
		}

		if i >= k-1 {
			l = append(l, nums[st[0]])
		}
	}
	return l
}
