package dp

/***

给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。

两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：

s = s1 + s2 + ... + sn
t = t1 + t2 + ... + tm
|n - m| <= 1
交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
注意：a + b 意味着字符串 a 和 b 连接。

输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出：true
示例 2：

输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
输出：false
示例 3：

输入：s1 = "", s2 = "", s3 = ""
输出：true

 */

/***
f[i][j] = f[i-1][j]     if s1[i] == s3[i+j+1]
f[i][j] = f[i][j-1]     if s2[j] == s3[i+j+1]
 */
func isInterleave(s1 string, s2 string, s3 string) bool {
	m := len(s1)
	n := len(s2)

	if n+m != len(s3) {
		return false
	}
	f := make([][]bool, m+1)

	// 太他妈的扯了， 这个初始化太难了。
	for i := 0; i <= m; i++ {
		f[i] = make([]bool, n+1)
		f[i][0] = s1[:i] == s3[:i]
	}
	for j := 1; j <= n; j++ {
		f[0][j] = s2[:j] == s3[:j]
	}

	f[0][0] = true

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s1[i] == s3[i+j+1] {
				f[i+1][j+1] = f[i+1][j+1] || f[i][j+1]
			}
			if s2[j] == s3[i+j+1] {
				f[i+1][j+1] = f[i+1][j+1] || f[i+1][j]
			}
		}
	}
	return f[m][n]
}
