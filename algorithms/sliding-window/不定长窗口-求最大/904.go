package weekly

/*

你正在探访一家农场，农场从左到右种植了一排果树。这些树用一个整数数组 fruits 表示，其中 fruits[i] 是第 i 棵树上的水果 种类 。

你想要尽可能多地收集水果。然而，农场的主人设定了一些严格的规矩，你必须按照要求采摘水果：

你只有 两个 篮子，并且每个篮子只能装 单一类型 的水果。每个篮子能够装的水果总量没有限制。
你可以选择任意一棵树开始采摘，你必须从 每棵 树（包括开始采摘的树）上 恰好摘一个水果 。采摘的水果应当符合篮子中的水果类型。每采摘一次，你将会向右移动到下一棵树，并继续采摘。
一旦你走到某棵树前，但水果不符合篮子的水果类型，那么就必须停止采摘。
给你一个整数数组 fruits ，返回你可以收集的水果的 最大 数目。



示例 1：

输入：fruits = [1,2,1]
输出：3
解释：可以采摘全部 3 棵树。
示例 2：

输入：fruits = [0,1,2,2]
输出：3
解释：可以采摘 [1,2,2] 这三棵树。
如果从第一棵树开始采摘，则只能采摘 [0,1] 这两棵树。
示例 3：

输入：fruits = [1,2,3,2,2]
输出：4
解释：可以采摘 [2,3,2,2] 这四棵树。
如果从第一棵树开始采摘，则只能采摘 [1,2] 这两棵树。
示例 4：

输入：fruits = [3,3,3,1,2,1,1,2,3,3,4]
输出：5
解释：可以采摘 [1,2,1,1,2] 这五棵树。
 */

// 好难的面试题，才1500多分。

// 给你提示了， sliding window 的问题

// 这个解法太复杂了， startpos 的引入，是能解决问题，但是还要更简单的题解。 看下面 正在的划窗实现吧！
//func totalFruit(fruits []int) int {
//	n := len(fruits)
//	startpos := make([]int, n)
//	l := 0
//	m := make(map[int]bool)
//	m[fruits[0]] = true
//	ans := 1
//
//	// 这就算是在枚举 right
//	for i := 1; i < n; i++ {
//		if fruits[i] == fruits[i-1] {
//			startpos[i] = startpos[i-1]
//		} else {
//			startpos[i] = i
//			if _, ok := m[fruits[i]]; !ok {
//				m[fruits[i]] = true
//			}
//			if len(m) == 3 {
//				l = startpos[i-1]
//				var k int
//				for k = range m {
//					if k != fruits[i] && k != fruits[i-1] {
//						break
//					}
//				}
//				delete(m, k)
//			}
//		}
//		ans = max(ans, i-l+1)
//	}
//	return ans
//}

// 来自宫水三叶。
func totalFruit(fruits []int) int {
	n := len(fruits)
	cnts := make([]int, n+1) // 统计 fruits[i] 种类的个数。 比保存 startpos 巧妙
	l := 0
	ans := 0
	total := 0

	// 这就算是在枚举 right
	for i := 0; i < n; i++ {
		cnts[fruits[i]]++
		if cnts[fruits[i]] == 1 {
			total++
		}
		for total > 2 {
			cnts[fruits[l]]--
			if cnts[fruits[l]] == 0 {
				total--
			}
			l++
		}
		ans = max(ans, i-l+1)
	}
	return ans
}
