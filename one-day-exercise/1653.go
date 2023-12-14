package one_day_exercise

/*
给你一个字符串 s ，它仅包含字符 'a' 和 'b'​​​​ 。
你可以删除 s 中任意数目的字符，使得 s 平衡 。当不存在下标对 (i,j) 满足 i < j ，且 s[i] = 'b' 的同时 s[j]= 'a' ，此时认为 s 是 平衡 的。
请你返回使 s 平衡 的 最少 删除次数。

例子：
ababaaaabbbbbaaababbbbbbaaabbaababbabbbbaabbbbaabbabbabaabbbababaa
*/

/*
	你花了足足2个小时，才写出来正确的前缀和的答案！ 再练！
	利用了前缀和的知识。 计算前缀和的代码容易出错。

https://leetcode.cn/problems/minimum-deletions-to-make-string-balanced/solutions/2149746/qian-hou-zhui-fen-jie-yi-zhang-tu-miao-d-dor2/
灵神给的答案，太绝了。 这道题，还可以用 DP来解答，尝试一下。
 */
func minimumDeletions(s string) int {
	n := len(s)
	// 前缀和, prea[i] 代表， 0.。i-1 之间出现了多少次 a
	prea := make([]int, n+1)
	preb := make([]int, n+1)

	ans := int(1e9)

	// 这样写，前缀和的计算，能好一点吗？
	for i := 0; i < n; i++ {
		prea[i+1] = prea[i]
		preb[i+1] = preb[i]
		if s[i] == 'a' {
			prea[i+1]++
		} else {
			preb[i+1]++
		}
	}

	for i := 0; i <= n; i++ {
		// 我们来枚举 a和b 分界的位置i，  i-1之前都是a, i 和 i 之后的位置都是 b。
		ans = min(ans, preb[i]+prea[n]-prea[i])
	}
	return ans
}

/*
真是好题啊，做几遍错几遍啊！ 这是标准的面试难度的真题。
 */
// 这个答案是错的。
func minimumDeletions_wrong(s string) int {
	// 删除 a , 再 b 出现之后，出现的 a 的次数。
	// 删除 b,  在最后一个 a 之前 所有 b 的出现次数。
	ans := 0
	bcnts := 0
	acnts := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			if bcnts > 0 {
				acnts++
			}
			ans = bcnts
		} else {
			bcnts++
		}
	}
	return min(acnts, ans)
}
