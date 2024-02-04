package difference

/***

给你一个二维整数数组 intervals ，其中 intervals[i] = [lefti, righti] 表示 闭 区间 [lefti, righti] 。

你需要将 intervals 划分为一个或者多个区间 组 ，每个区间 只 属于一个组，且同一个组中任意两个区间 不相交 。

请你返回 最少 需要划分成多少个组。

如果两个区间覆盖的范围有重叠（即至少有一个公共数字），那么我们称这两个区间是 相交 的。比方说区间 [1, 5] 和 [5, 8] 相交。



示例 1：

输入：intervals = [[5,10],[6,8],[1,5],[2,3],[1,10]]
输出：3
解释：我们可以将区间划分为如下的区间组：
- 第 1 组：[1, 5] ，[6, 8] 。
- 第 2 组：[2, 3] ，[5, 10] 。
- 第 3 组：[1, 10] 。
可以证明无法将区间划分为少于 3 个组。
示例 2：

输入：intervals = [[1,3],[5,6],[8,10],[11,13]]
输出：1
解释：所有区间互不相交，所以我们可以把它们全部放在一个组内。


提示：

1 <= intervals.length <= 10^5
intervals[i].length == 2
1 <= lefti <= righti <= 10^6

 */

/***

特别好的题目， 可以作为面试题目。
这个，处理区间的思路，是不是太巧妙了!
 */

func minGroups(intervals [][]int) int {
	n := 0 // 很搞笑，在我们没有 Loop intervals 之前，我们不知道 end 的值域， 只能从题目给的范围定义 1e6
	diffs := make([]int, int(1e6+2))

	for _, interval := range intervals {
		start, end := interval[0], interval[1]

		diffs[start] += 1
		//if end+1 < n {
		diffs[end+1] -= 1
		//}
		n = max(n, end)
	}
	needed := 0
	for i := 1; i <= n; i++ {
		diffs[i] += diffs[i-1]
		needed = max(needed, diffs[i])
	}
	return needed
}
