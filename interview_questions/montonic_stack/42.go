package montonic_stack

/***

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9

*/

/***
用单调栈来解决一下
wtf , 真的神奇。

我感觉的是， 开 left, 和 right 数组，比较容易记忆啊!
 */
func trap(height []int) int {
	n := len(height)
	left := make([]int, n)
	right := make([]int, n)
	for i := range left {
		left[i] = -1
		right[i] = n
	}
	ans := 0
	st := []int{}
	for i := 0; i < n; i++ {
		for len(st) > 0 && height[i] > height[st[len(st)-1]] {
			top := st[len(st)-1]
			right[top] = i
			if left[top] != - 1 {
				l, r := left[top], right[top]
				ans += (min(height[l], height[r]) - height[top]) * (r - l - 1)
			}
			// pop
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}

/****
前后缀，分解， 做出来的题， 这么简单吗？
多做啊， 熟练它
 */

func trap(height []int) int {
	n := len(height)
	left := make([]int, n)
	right := make([]int, n)

	for i := range left {
		left[i] = -1
		right[i] = n
	}

	// left wall
	lw := -1
	for i := 0; i < n; i++ {
		if lw == -1 || height[i] > height[lw] {
			lw = i
			left[i] = -1
		} else {
			left[i] = lw
		}
	}
	rw := -1
	for i := n - 1; i >= 0; i-- {
		if rw == -1 || height[i] > height[rw] {
			rw = i
			right[i] = -1
		} else {
			right[i] = rw
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if left[i] == -1 || right[i] == -1 {
			continue
		}
		ans += min(height[left[i]], height[right[i]]) - height[i]
	}
	return ans
}
