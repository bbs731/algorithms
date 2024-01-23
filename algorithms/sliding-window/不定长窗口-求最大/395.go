package weekly

/***

给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。

如果不存在这样的子字符串，则返回 0。


示例 1：

输入：s = "aaabb", k = 3
输出：3
解释：最长子串为 "aaa" ，其中 'a' 重复了 3 次。
示例 2：

输入：s = "ababbc", k = 2
输出：5
解释：最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。

 */

// 好难的题， 没发现套路，还得挂！
// 理论升华了：
//https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters/solutions/624045/xiang-jie-mei-ju-shuang-zhi-zhen-jie-fa-50ri1/

func longestSubstring(s string, k int) int {
	n := len(s)
	cnts := make(map[int32]int, n)
	for _, c := range s {
		cnts[c]++
	}
	ans := 0

	// 枚举的力量啊， 通过枚举字符集的大小， 让区间重新具有了二段性。 太美妙了！
	for t := 1; t <= 26; t++ {
		left := 0
		wc := make(map[int32]int, n)
		tobf := make(map[int32]struct{}, n)

		// 枚举右端点
		for right, c := range s {
			wc[c]++
			if wc[c] >= k {
				delete(tobf, c)
			} else {
				tobf[c] = struct{}{}
			}
			for len(wc) > t {
				lc := int32(s[left])
				wc[lc]--
				if wc[lc] < k {
					tobf[lc] = struct{}{}
				}
				if wc[lc] == 0 {
					delete(wc, lc)
					delete(tobf, lc)
				}
				left++
			}

			if len(tobf) == 0 {
				ans = max(ans, right-left+1)
			}
		}
	}
	return ans
}
