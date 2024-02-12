package bits_operation

import "math/bits"

/***
给你两个整数 left 和 right ，表示区间 [left, right] ，返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）。

示例 1：
输入：left = 5, right = 7
输出：4
示例 2：

输入：left = 0, right = 0
输出：0
示例 3：

输入：left = 1, right = 2147483647
输出：0


提示：
0 <= left <= right <= 2^31 - 1
 */

/***
https://leetcode.cn/problems/bitwise-and-of-numbers-range/solutions/538550/golang-yi-xing-suan-fa-by-endlesscheng-iw6y/

大哥是按位 与(AND) 操作， 不是 异或 XOR
 */

func rangeBitwiseAnd(m, n int) int {
	return m &^ (1<<bits.Len(uint(m^n)) - 1)
}
