package monotonic_stack

/***

给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。

输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10



输入： heights = [2,4]
输出： 4

1 <= heights.length <=10^5
0 <= heights[i] <= 10^4

 */

/***

 宫三的题解：
 https://leetcode.cn/problems/largest-rectangle-in-histogram/solutions/1856193/by-ac_oier-i470/

 枚举每个柱子， 找到柱子的左右边界， 就是，左右第一个比柱子小的坐标。（这个不是单调栈的模版吗？）

 能想到，这么解的都是天才。
 */


func largestRectangleArea(heights []int) int {
	n :=len(heights)
	left := make([]int, n)
	right := make([]int, n)
	for i := range right {
		right[i]= n
	}

	st := []int{-1}
	for i, v := range heights {
		for len(st)	 > 1 && v < heights[st[len(st)-1]] {
			right[st[len(st)-1]] = i
			// pop stack
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}

	ans := 0
	for i, v := range heights {
		ans = max(ans, (right[i] - left[i]-1)*v)
	}
	return ans
}
