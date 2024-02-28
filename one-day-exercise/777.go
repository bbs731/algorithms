package one_day_exercise

/***
在一个由 'L' , 'R' 和 'X' 三个字符组成的字符串（例如"RXXLRXRXL"）中进行移动操作。一次移动操作指用一个"LX"替换一个"XL"，或者用一个"XR"替换一个"RX"。现给定起始字符串start和结束字符串end，请编写代码，当且仅当存在一系列移动操作使得start可以转换成end时， 返回True。



示例 :

输入: start = "RXXLRXRXL", end = "XRLXXRRLX"
输出: True
解释:
我们可以通过以下几步将start转换成end:
RXXLRXRXL ->
XRXLRXRXL ->
XRLXRXRXL ->
XRLXXRRXL ->
XRLXXRRLX


提示：

1 <= len(start) = len(end) <= 10000。
start和end中的字符串仅限于'L', 'R'和'X'。
 */

/****

这种怎么想的？  L只能往左移动， R 只能向右。

 */

func canTransform(start string, end string) bool {
	n := len(start)
	i, j := 0, 0

	for true {
		for ; i < n && start[i] == 'X'; i++ {
		}
		for ; j < n && end[j] == 'X'; j++ {
		}
		if i == n || j == n {
			return i == j
		}
		if start[i] != end[j] {
			return false
		}
		if start[i] == 'L' && i < j {
			return false
		}
		if start[i] == 'R' && i > j {
			return false
		}
		i++
		j++
	}
	return false
}

/****
思路错误！
 */
//func canTransform(start string, end string) bool {
//	n := len(start)
//	if len(end) != n {
//		return false
//	}
//	dp := make([]bool, n+1)
//	dp[n] = true
//	if start[n-1] == end[n-1] {
//		dp[n-1] = true
//	}
//
//	for i := n - 2; i >= 0; i-- {
//		if start[i] == end[i] {
//			dp[i] = dp[i+1]
//		} else {
//			if (start[i] == 'R' && start[i+1] == 'X') || (start[i] == 'X' && start[i+1] == 'L') {
//				if start[i] == end[i+1] && end[i] == start[i+1] {
//					dp[i] = dp[i+2]
//				} else if end[i+1] == 'X' {
//					dp[i+1] = dp[i+2]
//				}
//			}
//		}
//	}
//	return dp[0]
//}
