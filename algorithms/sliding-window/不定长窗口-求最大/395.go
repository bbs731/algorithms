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

bbaaacbd

k=5
aaaaaaaaabbbcccccddddd

k = 10
aaaaaaaaabbbcccccddddd
zzzzzzzzzzaaaaaaaaabbbbbbbbhbhbhbhbhbhbhicbcbcibcbccccccccccbbbbbbbbaaaaaaaaafffaahhhhhiaahiiiiiiiiifeeeeeeeeee
 */

// 好难的题， 没发现套路，还得挂！

func longestSubstring(s string, k int) int {
	n := len(s)
	cnts := make(map[int32]int, n)
	for _, c := range s {
		cnts[c]++
	}
	left := 0
	ans := 0

	wc := make(map[int32]int, n)
	tobf := make(map[int32]struct{}, n)

	// 枚举右端点
	for right, c := range s {
		if cnts[c] < k && len(wc) > 0 {
			// 这里会漏掉可能的答案。
			//for k, v := range wc {
			//	cnts[k] -= v
			//}
			repeat := 1
			for j := left; j <= right-1; j++ {
				if j > left {
					if s[j] == s[j-1] {
						repeat++
						if repeat >= k {
							ans = max(ans, repeat)
						}
					} else {
						repeat = 1
					}
				}
				cnts[int32(s[j])]--
			}

			left = right + 1
			wc = make(map[int32]int, n)
			tobf = make(map[int32]struct{}, n)
		} else {
			wc[c]++
			if wc[c] >= k {
				delete(tobf, c)
				if len(tobf) == 0 {
					ans = max(ans, right-left+1)
				}
			} else {
				tobf[c] = struct{}{}
			}
		}
	}
	return ans
}
