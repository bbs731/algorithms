package _400_1600

/*

给你一个二进制串 s  （一个只包含 0 和 1 的字符串），我们可以将 s 分割成 3 个 非空 字符串 s1, s2, s3 （s1 + s2 + s3 = s）。

请你返回分割 s 的方案数，满足 s1，s2 和 s3 中字符 '1' 的数目相同。

由于答案可能很大，请将它对 10^9 + 7 取余后返回。



示例 1：

输入：s = "10101"
输出：4
解释：总共有 4 种方法将 s 分割成含有 '1' 数目相同的三个子字符串。
"1|010|1"
"1|01|01"
"10|10|1"
"10|1|01"
示例 2：

输入：s = "1001"
输出：0
示例 3：

输入：s = "0000"
输出：3
解释：总共有 3 种分割 s 的方法。
"0|0|00"
"0|00|0"
"00|0|0"
示例 4：

输入：s = "100100010100110"
输出：12


提示：

s[i] == '0' 或者 s[i] == '1'
3 <= s.length <= 10^5

 */

/*

难度分数：  1591

你找到了规律， 但是这种题， 一旦规律找到， 一点难度都没有，只是统计。和考虑边界情况。
这种题，叫做模拟。 少做吧， 太水了，涨不了知识！
 */
func numWays(s string) int {
	mod := int(1e9) + 7
	n := len(s)
	nones := 0
	for i := range s {
		if s[i] == '1' {
			nones++
		}
	}

	if nones == 0 {
		// n-1 选出  2个位置  C(n-1, 2)
		return (n - 1) * (n - 2) / 2 % mod
	}

	if nones%3 != 0 {
		return 0
	}

	third := nones / 3
	thirdpos1 := 0
	thirdpos2 := 0
	cnt := 0
	lzeros1 := 0
	lzeros2 := 0
	for i := range s {
		if s[i] == '1' {
			cnt++
			if thirdpos1 != 0 {
				thirdpos1 = 0
			}
			if thirdpos2 != 0 {
				thirdpos2 = 0
			}
		} else {
			if thirdpos1 != 0 {
				lzeros1++
			}
			if thirdpos2 != 0 {
				lzeros2++
			}
		}
		if cnt == third {
			thirdpos1 = 100
		}
		if cnt == 2*third {
			thirdpos2 = 100
		}
	}
	return (lzeros1 + 1) * (lzeros2 + 1) % mod
}
