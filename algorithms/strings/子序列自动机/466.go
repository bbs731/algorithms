package subsequenceAutomation

/***
定义 str = [s, n] 表示 str 由 n 个字符串 s 连接构成。

例如，str == ["abc", 3] =="abcabcabc" 。
如果可以从 s2 中删除某些字符使其变为 s1，则称字符串 s1 可以从字符串 s2 获得。

例如，根据定义，s1 = "abc" 可以从 s2 = "abdbec" 获得，仅需要删除加粗且用斜体标识的字符。
现在给你两个字符串 s1 和 s2 和两个整数 n1 和 n2 。由此构造得到两个字符串，其中 str1 = [s1, n1]、str2 = [s2, n2] 。

请你找出一个最大整数 m ，以满足 str = [str2, m] 可以从 str1 获得。



示例 1：

输入：s1 = "acb", n1 = 4, s2 = "ab", n2 = 2
输出：2
示例 2：

输入：s1 = "acb", n1 = 1, s2 = "acb", n2 = 1
输出：1


提示：

1 <= s1.length, s2.length <= 100
s1 和 s2 由小写英文字母组成
1 <= n1, n2 <= 10^6

 */

/***
灵神的版本， 处理， pattern repeat 的时候，有问题。 还是用自己的版本吧。
 */
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	// build nxt
	// nxt[i][j] 表示在i的右侧, 字符 j (第一次出现）的最近位置
	s := s1
	pos := [26]int{}
	for i := range pos {
		pos[i] = len(s) + 1
	}
	nxt := make([][26]int, len(s)+1)
	for i := len(s) - 1; i >= 0; i-- {
		nxt[i+1] = pos
		pos[s[i]-'a'] = i + 1
	}
	nxt[0] = pos // 灵神的版本 是没有保存最后一个 pos的, 处理的方法，留在了 match function 去比较一下 t[0] == s[0]

	match := func(t string) int {
		cnts1, cnts2 := 0, 0
		i, j := 0, 0
		for true {
			for ; j < len(s2); j++ {
				i = nxt[i][s2[j]-'a']
				if i == len(s)+1 {
					i = 0
					cnts1++
					if cnts1 >= n1 {
						return cnts2
					}
					break
				}
			}
			if j == len(s2) {
				cnts2++
				j = 0
			}
		}
		return 0
	}
	l := match(s2)
	return l / n2
}

/***
灵神的板子，在处理， s1 可以是 n 个 s1 repeated 字符串的时候，会有问题。 用自己的板子把。
 */
//
//func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
//	// build nxt
//	// nxt[i][j] 表示在i的右侧, 字符 j (第一次出现）的最近位置
//	s := s1
//	pos := [26]int{}
//	for i := range pos {
//		pos[i] = len(s) // 初始化， mark end, 在 match 的时候会判断使用
//	}
//	nxt := make([][26]int, len(s))
//
//	for i := len(s) - 1; i >= 0; i-- {
//		nxt[i] = pos
//		pos[s[i]-'a'] = i // 其实是为了 nxt[i-1] 准备的 pos 值。
//		// 这样写边界处理的很清晰， 但是有一个问题就是 nxt[0] = pos[s[1]-'a'], 我们发现 没有 nxt[-1] 也就是说
//		// pos[s[0]-'a'] 的这个 match 关系我们没有保存，是丢掉了的。针对这个问题，就需要在 match 里对 s[0] 特殊处理。
//	}
//
//	match := func(t string) int {
//		i, j := 0, 0
//		cnts1, cnts2 := 0, 0
//		for true {
//			// 因为我们在 build nxt 的时候，丢弃了 pos[s[0]-'a'] 这个信息，所以需要在这里，特殊处理一下。
//			if i == 0 {
//				if t[j] == s[0] {
//					j = j + 1 // t[0] 匹配。
//				}
//			}
//			for ; j < len(t); j++ {
//				i = nxt[i][t[j]-'a']
//				if i == len(s) {
//					// 很多题，可以在这里处理，想到到的结果。
//					cnts1++
//					i = 0
//					if cnts1 >= n1 {
//						return cnts2
//					}
//					break
//				}
//			}
//
//			if j == len(t) {
//				//found a match
//				cnts2++
//				j = 0
//			}
//		}
//		return 0
//	}
//	l := match(s2)
//	return l / n2
//}
