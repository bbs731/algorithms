package dp

/***
给你一个输入字符串 (s) 和一个字符模式 (p) ，请你实现一个支持 '?' 和 '*' 匹配规则的通配符匹配：
'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符序列（包括空字符序列）。
判定匹配成功的充要条件是：字符模式必须能够 完全匹配 输入字符串（而不是部分匹配）。


示例 1：

输入：s = "aa", p = "a"
输出：false
解释："a" 无法匹配 "aa" 整个字符串。
示例 2：

输入：s = "aa", p = "*"
输出：true
解释：'*' 可以匹配任意字符串。
示例 3：

输入：s = "cb", p = "?a"
输出：false
解释：'?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。


提示：
0 <= s.length, p.length <= 2000
 */

/***
这个，竟然是一道 DP 的问题， 怎么也想不到啊？  我靠， 见了鬼了， 这么难吗?

f[i][j] =    f[i-1][j-1]  if s[i] == p[j] for p[j] == '?'
		=   or(f[k][j-1])    if p[j] == '*'  for all k = 0..... i
	   // 枚举 这个 k 有更简洁的写法。  看官方的题解， 但是不知道为什么？ 可以， 如何证明呢？ // 我才这个巧妙了， 保留了 * 去匹配更小的 i, 不用枚举 i 啊， 因为 * 还保留这。 哎，真厉害！
*/

func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)

	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true

	// 就是初始化有问题！ 这里。
	for i := 1; i <= m; i++ {
		if p[i-1] == '*' {
			f[0][i] = true
		} else {
			break
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if p[j] == '*' {
				for k := 0; k <= i+1; k++ {
					if f[k][j] {
						f[i+1][j+1] = true
						break
					}
				}
			} else {
				if s[i] == p[j] || p[j] == '?' {
					f[i+1][j+1] = f[i][j]
				}
			}
		}
	}
	return f[n][m]
}

func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)

	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true

	// 就是初始化有问题！ 这里。
	for i := 1; i <= m; i++ {
		if p[i-1] == '*' {
			f[0][i] = true
		} else {
			break
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if p[j] == '*' {
				f[i+1][j+1] = f[i+1][j] || f[i][j+1]
			} else {
				if s[i] == p[j] || p[j] == '?' {
					f[i+1][j+1] = f[i][j]
				}
			}
		}
	}
	return f[n][m]
}
