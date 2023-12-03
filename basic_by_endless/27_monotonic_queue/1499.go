package mono_queue



/*
这道题， 我们调整了一下，顺序， 先出队，再入队。

另外求的最大值为，  yi + yj + xj-xi  = yj +xj + (yi-xi)   既然yj+xj 是定值（正在处理 point j) 那么 单调队列中我们应该根据 yi-xi 排序。


https://leetcode.cn/problems/max-value-of-equation/solutions/2352457/on-dan-diao-dui-lie-fu-ti-dan-pythonjava-hhrr/
看看灵神的讲解，思路是一样的，代码更加的优美
 */
//[[-19,-12],[-5,-18],[2,-2],[10,3],[11,-3],[13,17]] k =13
func findMaxValueOfEquation(points [][]int, k int) int {
	q := []int{}
	ans := -int(1e10)

	for i, point := range points {

		//出队
		for len(q) > 0 && point[0] - points[q[0]][0] > k{
			q = q[1:]
		}

		if len(q) > 0 { //题目保证了，至少有一个解
			ans = max(ans, points[q[0]][1] + point[1] + point[0] - points[q[0]][0])
		}

		// 入队。
		for len(q) > 0 && points[q[len(q)-1]][1] - points[q[len(q)-1]][0] < point[1] -point[0] {
			q = q[:len(q)-1]
		}
		q = append(q, i)

	}

	return ans
}
