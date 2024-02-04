package difference

/***

http://leetcode.cn/circle/discuss/FfMCgb
差分数组的定义：

d[0] = a[0]   // 这个特别重要
d[i] = a[i] - a[i-1]  for i >=1

性质1： 从左到右累加d中的元素，可以得到数组a
性质2： 把区间造作，变成单点操作， 把 a[left] .....  a[right] 之间的数都加 x 等价于
d[left] += x
d[right+1] -= x    (if right+1 <n) 这里特别 right+1=n 我们只需要 d[left] +=x。 因为最后恢复数组 a 和 d[n] 没有关系。n 数组 a 长度

 */

// 你有一个长为 n 的数组 a，一开始所有元素均为 0。
// 给定一些区间操作，其中 queries[i] = [left, right, x]，
// 你需要把子数组 a[left], a[left+1], ... a[right] 都加上 x。
// 返回所有操作执行完后的数组 a。
func solve(n int, queries [][]int) []int {
	// 模板代码
	diff := make([]int, n) // 差分数组
	for _, q := range queries {
		left, right, x := q[0], q[1], q[2]
		diff[left] += x
		if right+1 < n {
			diff[right+1] -= x
		}
	}
	for i := 1; i < n; i++ {
		diff[i] += diff[i-1] // 直接在差分数组上复原数组 a
	}
	return diff
}
