package sliding_window

/***
给定一个包含非负整数的数组 nums ，返回其中可以组成三角形三条边的三元组个数。

示例 1:
输入: nums = [2,2,3,4]
输出: 3
解释:有效的组合是:
2,3,4 (使用第一个 2)
2,3,4 (使用第二个 2)
2,2,3
示例 2:

输入: nums = [4,2,3,4]
输出: 4


提示:

1 <= nums.length <= 1000
0 <= nums[i] <= 1000

 */

/*

这道题，太赞了！

这个写法，的时间复杂度是 O(n^2) 不是 n^3
 */
func triangleNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	ans := 0

	for i := 0; i < n; i++ {
		a := nums[i]
		l, r := i+1, n-1
		for l < r {
			if a+nums[l] > nums[r] {
				ans += r - l
				r--
			} else if a+nums[l] <= nums[r] {
				// wo kao 这个技巧学到了， r-- or l++  先处理 r-- 的情况。
				for k := r - 1; k > l; k-- {
					if a+nums[l] > nums[k] {
						ans += k - l
						break
					}
				}
				l++
			}
		}
	}
	return ans
}
