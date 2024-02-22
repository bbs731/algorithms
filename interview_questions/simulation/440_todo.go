package simulation

/***

https://leetcode.cn/problems/k-th-smallest-in-lexicographical-order/solutions/1358635/zi-dian-xu-de-di-kxiao-shu-zi-by-leetcod-bfy0/

官网的逻辑，真是太清楚了。

自己做的话， 怎么才能想得到呢？ 这是一道模拟的题目， 需要准确的知道，模拟的过程和步骤。


说实话， 这个，现场想的话，够呛， 就靠背了吗？
*/

func getCurrent(cur, n int) int {
	first, last := cur, cur
	step := 0
	for first <= n {
		step += min(last, n) - first + 1
		first *= 10
		last = last*10 + 9
	}
	return step
}

func findKthNumber(n int, k int) int {
	cur := 1
	k--
	for k > 0 {
		if step := getCurrent(cur, n); step <= k {
			cur++
			k -= step
		} else {
			cur *= 10
			k--
		}
	}
	return cur
}
