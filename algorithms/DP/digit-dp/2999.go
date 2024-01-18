package digit_dp

import (
	"strconv"
	"strings"
)

/*

给你三个整数 start ，finish 和 limit 。同时给你一个下标从 0 开始的字符串 s ，表示一个 正 整数。

如果一个 正 整数 x 末尾部分是 s （换句话说，s 是 x 的 后缀），且 x 中的每个数位至多是 limit ，那么我们称 x 是 强大的 。

请你返回区间 [start..finish] 内强大整数的 总数目 。

如果一个字符串 x 是 y 中某个下标开始（包括 0 ），到下标为 y.length - 1 结束的子字符串，那么我们称 x 是 y 的一个后缀。比方说，25 是 5125 的一个后缀，但不是 512 的后缀。



示例 1：

输入：start = 1, finish = 6000, limit = 4, s = "124"
输出：5
解释：区间 [1..6000] 内的强大数字为 124 ，1124 ，2124 ，3124 和 4124 。这些整数的各个数位都 <= 4 且 "124" 是它们的后缀。注意 5124 不是强大整数，因为第一个数位 5 大于 4 。
这个区间内总共只有这 5 个强大整数。
示例 2：

输入：start = 15, finish = 215, limit = 6, s = "10"
输出：2
解释：区间 [15..215] 内的强大整数为 110 和 210 。这些整数的各个数位都 <= 6 且 "10" 是它们的后缀。
这个区间总共只有这 2 个强大整数。
示例 3：

输入：start = 1000, finish = 2000, limit = 4, s = "3000"
输出：0
解释：区间 [1000..2000] 内的整数都小于 3000 ，所以 "3000" 不可能是这个区间内任何整数的后缀。

 */

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	highS := strconv.Itoa(int(finish))
	lowS := strconv.Itoa(int(start))
	n := len(highS)
	lowS = strings.Repeat("0", n-len(lowS)) + lowS
	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1
	}

	var dfs func(int, bool, bool) int
	dfs = func(i int, limitLow, limitHigh bool) (res int) {
		if i == n {
			return 1
		}

		if !limitLow && !limitHigh {
			if dp[i] != -1 {
				return dp[i]
			}
		}
		if i >= n-len(s) {
			if limitLow && lowS[i] > s[i+len(s)-n] {
				return
			}
			if limitHigh && highS[i] < s[i+len(s)-n] {
				return
			}
			res += dfs(i+1, limitLow && lowS[i] == s[i+len(s)-n], limitHigh && highS[i] == s[i+len(s)-n])
		} else {
			lo := 0
			if limitLow {
				lo = int(lowS[i] - '0')
			}
			hi := 9
			if limitHigh {
				hi = int(highS[i] - '0')
			}
			for d := lo; d <= min(hi, limit); d++ {
				res += dfs(i+1, limitLow && d == lo, limitHigh && d == hi)
			}
		}

		if !limitLow && !limitHigh {
			dp[i] = res
		}
		return
	}

	return int64(dfs(0, true, true))
}
