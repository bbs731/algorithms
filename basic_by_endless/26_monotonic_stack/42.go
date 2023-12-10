package monotonic_stack

/*
灵神用单调栈的解法：
 https://leetcode.cn/problems/trapping-rain-water/solution/zuo-liao-nbian-huan-bu-hui-yi-ge-shi-pin-ukwm/
 */
func trap(height []int) int {
	ans := 0
	st := make([]int, 0)

	for i, t := range height {
		for len(st) > 0 && t >= height[st[len(st)-1]] {
			top := st[len(st)-1]
			//pop top
			st = st[:len(st)-1]
			if len(st) > 0 {
				w := i - st[len(st)-1]-1
				h := min(height[st[len(st)-1]], t) - height[top]
				ans += w * h
			}
		}
		st = append(st, i)
	}
	return ans
}
