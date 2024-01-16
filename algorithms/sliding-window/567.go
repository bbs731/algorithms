package sliding_window

/***
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。

换句话说，s1 的排列之一是 s2 的 子串 。



示例 1：

输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").
示例 2：

输入：s1= "ab" s2 = "eidboaoo"
输出：false


1 <= s1.length, s2.length <= 10^4
s1 和 s2 仅包含小写字母

 */

func checkPermutation(s1, s2 string) bool {
	m := make(map[int32]bool)
	m2 := make(map[int32]bool)
	for _, s := range s1 {
		m[s] = true
	}
	for _, s := range s2 {
		if _, ok := m[s]; !ok {
			return false
		}
		m2[s] = true
	}
	if len(m) != len(m2) {
		return false
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	n1 := len(s1)
	n2 := len(s2)

	if n1 > n2 {
		return false
	}
	sum1 := 0
	sum2 := 0
	for i := 0; i < n1; i++ {
		sum1 += int(s1[i] - 'a')
		sum2 += int(s2[i] - 'a')
	}
	sum2 -= int(s2[n1-1] - 'a')
	for i := n1 - 1; i < n2; i++ {
		sum2 += int(s2[i] - 'a')
		if sum1 == sum2 {
			if checkPermutation(s1, s2[i-n1+1:i+1]) {
				return true
			}
		}
		sum2 -= int(s2[i-n1+1] - 'a')
	}
	return false
}

// 官方题解： 第一次感觉到， 官方的题解，还能这么优雅！ 哎！
// 牛！
func checkInclusion(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	var cnt1, cnt2 [26]int
	for i, ch := range s1 {
		cnt1[ch-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}
	for i := n; i < m; i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-n]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}
