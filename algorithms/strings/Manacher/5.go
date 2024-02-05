package Manacher

/***
给你一个字符串 s，找到 s 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。



示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"


提示：

1 <= s.length <= 1000
s 仅由数字和英文字母组成

 */

func longestPalindrome(s string) string {
	manacher := func(s string) []int {
		//把 s 改造为字符串 t
		t := append(make([]byte, 0, len(s)*2+3), '^')
		for _, c := range s {
			t = append(t, '#', byte(c))
		}
		t = append(t, '#', '$')

		// 定义一个奇回文串的回文半径=(长度+1)/2，即保留回文中心，去掉一侧后的剩余字符串的长度 (halfLen 的值)
		// 完整的定义（需要指定中心位置)： halfLen[i] 表示在 t 上的以 t[i] 为回文中心的最长回文子串的回文半径
		// 即 [i-halfLen[i]+1,i+halfLen[i]-1] 是 t 上的一个回文子串
		// halfLen 数组，即为 manacher 算法的所求。

		halfLen := make([]int, len(t)-2)
		halfLen[1] = 1

		// 一般的实现会用 l, r 表示回文的左右端点， 灵神的实现，创新的使用了 mid 和 halfLen 有来简化实现,避免考虑太多 +1 -1 的地方。
		// 定义：mid 为该回文子串的中心位置，定义 r 为 r=mid+halfLen[mid]
		// r 表示当前右边界下标最大的回文子串的右边界下标+1.  [mid-halfLen[mid]+1,  mid+halfLen[mid]-1] 这个区间是真正的回文字符串。

		for i, mid, r := 2, 1, 0; i < len(halfLen); i++ {
			hl := 1
			if i < r {
				// i 关于 mid 的对称位置 i' = mid*2 -i     由: mid-i = i-mid 可以得出
				// 若以 i' 为中心的最长回文子串范围超出了以 mid 为中心的回文串的范围（即 i+halfLen[i'] >= r）
				// 则 halfLen[i] 应先初始化为已知的回文半径 r-i，然后再继续暴力匹配
				// 否则 halfLen[i] 与 halfLen[i'] 相等
				hl = min(halfLen[mid*2-i], r-i) // manacher 算法是线性的关键就是因为这条优化
			}
			// 暴力扩展
			// 算法的复杂度取决于这部分执行的次数
			// 由于扩展之后 r 必然会更新（右移），且扩展的的次数就是 r 右移的次数
			// 因此算法的复杂度和 t 的长度成正比
			for t[i-hl] == t[i+hl] {
				hl++
			}

			if i+hl > r {
				mid, r = i, i+hl
			}
			halfLen[i] = hl
		}

		return halfLen
	}

	ans := string(s[0])
	ansl := 1
	halfLen := manacher(s)
	for i := 2; i < len(halfLen); i++ {
		if halfLen[i]-1 > ansl {
			ansl = halfLen[i] - 1
			// [(i-hl)/2, (i+hl)/2-2]  // 对应的在 s 中的回文串。
			ans = s[(i-halfLen[i])/2 : (i+halfLen[i])/2-1]
		}
	}
	return ans
}
