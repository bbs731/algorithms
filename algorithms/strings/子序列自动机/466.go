package subsequenceAutomation

import "fmt"

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
		pos[i] = len(s) + 2
	}
	nxt := make([][26]int, len(s)+1)
	for i := len(s) - 1; i >= 0; i-- {
		nxt[i+1] = pos
		pos[s[i]-'a'] = i + 1
	}
	nxt[0] = pos // 灵神的版本 是没有保存最后一个 pos的, 处理的方法，留在了 match function 去比较一下 t[0] == s[0]

	match := func(t string) int{
		cnts1, cnts2 := 0, 0
		i, j := 0, 0
		for true {
			for ; j < len(s2); j++ {
				i = nxt[i][s2[j]-'a']
				if i == len(s)+2 {
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
				j=0
			}
		}
		return 0
	}
	l :=match(s2)
	return l/n2
}


func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	// build nxt
	// nxt[i][j] 表示在i的右侧, 字符 j (第一次出现）的最近位置
	s := s1
	pos := [26]int{}
	for i := range pos {
		pos[i] = len(s) + 2
	}
	nxt := make([][26]int, len(s)+1)
	for i := len(s) - 1; i >= 0; i-- {
		nxt[i+1] = pos
		pos[s[i]-'a'] = i + 1
	}
	nxt[0] = pos // 灵神的版本 是没有保存最后一个 pos的, 处理的方法，留在了 match function 去比较一下 t[0] == s[0]

	match := func(t string, m int) bool {
		cnts1, cnts2 := 0, 0
		i, j := 0, 0
		for true {
			for ; j < len(s2); j++ {
				i = nxt[i][s2[j]-'a']
				if i == len(s)+2 {
					i = 0
					cnts1++
					if cnts1 >= n1 {
						return false
					}
					break
				}
			}
			if j == len(s2) {
				cnts2++
				if cnts2 >= m {
					return true
				}
				j=0
			}
		}
		return false
	}

	l, r := 0, len(s1)*n1/len(s2)+1
	fmt.Println(l,r)
	for l+1 < r {
		mid := (l + r) >> 1
		if match(s2, mid*n2) {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}

