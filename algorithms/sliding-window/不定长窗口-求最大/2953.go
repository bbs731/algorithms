package weekly

import "fmt"

/***


给你一个字符串 word 和一个整数 k 。

如果 word 的一个子字符串 s 满足以下条件，我们称它是 完全字符串：

s 中每个字符 恰好 出现 k 次。
相邻字符在字母表中的顺序 至多 相差 2 。也就是说，s 中两个相邻字符 c1 和 c2 ，它们在字母表中的位置相差 至多 为 2 。
请你返回 word 中 完全 子字符串的数目。

子字符串 指的是一个字符串中一段连续 非空 的字符序列。



示例 1：

输入：word = "igigee", k = 2
输出：3
解释：完全子字符串需要满足每个字符恰好出现 2 次，且相邻字符相差至多为 2 ：igigee, igigee, igigee 。
示例 2：

输入：word = "aaabbbccc", k = 3
输出：6
解释：完全子字符串需要满足每个字符恰好出现 3 次，且相邻字符相差至多为 2 ：aaabbbccc, aaabbbccc, aaabbbccc, aaabbbccc, aaabbbccc, aaabbbccc 。

灵神的题解：
https://leetcode.cn/problems/count-complete-substrings/solutions/2551743/bao-li-hua-chuang-mei-ju-chuang-kou-nei-31j5m/
 */

// 这是 OI 选手，变态的思维吗？ 再一次打脸了， 东拼西凑是做不出来好结果的！ 枚举的力量，充分利用题目的信息！和 395一样刷新了认知了
func f(s string, k int) (res int) {
	for m := 1; m <= 26 && k*m <= len(s); m++ {
		cnt := [26]int{}
		check := func() {
			for i := range cnt {
				if cnt[i] > 0 && cnt[i] != k {
					return
				}
			}
			res++
		}

		// 这里就直接编程  k*m  fix window 的划窗问题了。
		for right, in := range s {
			cnt[in-'a']++
			if left := right + 1 - k*m; left >= 0 {
				check()
				cnt[s[left]-'a']--
			}
		}
	}
	return
}

func countCompleteSubstrings(word string, k int) (ans int) {
	for i, n := 0, len(word); i < n; {
		st := i
		for i++; i < n && abs(int(word[i])-int(word[i-1])) <= 2; i++ {
		}
		ans += f(word[st:i], k)
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	};
	return x
}

func abs(a, b int32) int32 {
	if a > b {
		return a - b
	}
	return b - a
}

// 做不出来， 无论，如何， 都有重复，或者漏掉的情况。
// 不具备有单调性， 所以是解不出来的。 仿照 395 题来制造 二段性。
func countCompleteSubstrings(word string, k int) int {
	n := len(word)
	ans := 0
	left := 0
	//counter := 0
	//pos := make(map[pair]struct{})

	wc := make(map[int32]int, n)
	tof := make(map[int32]struct{}, n) // to be fill, char has not been accumulated to k
	for right, w := range word {
		if right > 0 && abs(w, int32(word[right-1])) > 2 {
			left = right
			wc = make(map[int32]int, n)
			tof = make(map[int32]struct{}, n)
		}
		wc[w]++
		if wc[w] < k {
			tof[w] = struct{}{}
		} else if wc[w] == k {
			delete(tof, w)
			if len(tof) == 0 {
				// 用 counter 还是 +1 都有测试的反例， 不如记录 left, right 然后统计数目。
				//counter++
				ans += 1
				//pos[pair{left, right}] = struct{}{}
			}
			// 这里会漏掉解。以 right 为右端点， 但是左端点 > left 这段区间，可能存在额外的解。
			if right != n-1 {
				tmp := countCompleteSubstrings(word[left+1:right+1], k)
				ans += tmp
				fmt.Println(left+1, right, tmp)
			}
		} else {
			//wc[w] = k + 1
			for left < right {
				lc := int32(word[left])
				if lc == w {
					wc[lc]--
					left++
					break
				}

				//if wc[lc] == k {
				//	counter--
				//}

				wc[lc]--
				if wc[lc] == 0 {
					delete(tof, lc)
				}
				left++
			}

			// after shift left window
			// double check whether we have a valid answer here
			if len(tof) == 0 {
				ans += 1
				//pos[pair{left, right}] = struct{}{}
			}

			if right != n-1 && right > left {
				tmp := countCompleteSubstrings(word[left+1:right+1], k)
				ans += tmp
				fmt.Println(left+1, right, tmp)
			}
		}
	}

	// 处理最后一段
	for left < n-1 {
		lc := int32(word[left])
		wc[lc]--
		// 0 <= wc[lc] < k
		tof[lc] = struct{}{}
		if wc[lc] == 0 {
			delete(tof, lc)
		}

		if len(tof) == 0 && len(wc) > 0 {
			ans += 1
			//pos[pair{left, n - 1}] = struct{}{}
		}
		left++
	}

	//return len(pos)
	return ans
}
