package digit_dp

import (
	"strconv"
	"strings"
)

/***

copy from 灵茶山
https://leetcode.cn/problems/count-the-number-of-powerful-integers/solutions/2595149/shu-wei-dp-shang-xia-jie-mo-ban-fu-ti-da-h6ci/


Digit-DP 有两个模版
1.0 版本，没有下界的约束
2.0 版本，有下界的约束。


// 时间复杂度 ???
// Digig DP template 1.0 版面。
https://www.bilibili.com/video/BV1rS4y1s721/
https://leetcode.cn/problems/count-special-integers/solutions/1746956/shu-wei-dp-mo-ban-by-endlesscheng-xtgx/

***/

func digit_dp_template(n int) int {

	// step 1
	// convert to string 如果是 10进制的题目
	// m := len(s)
	// 计算二进制的长度， 如果是二进制题目。
	//m := bits.Len(uint(n))
	//memo := make([][]int, m)

	s := strconv.Itoa(n)
	m := len(s)
	memo := make([][1 << 10]int, m)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}

	var f func(int, int, bool, bool) int
	// mask 不是一定的， 有的题目不需要 mask, 需要 cnt1 譬如 count 1 出现的次数。
	f = func(i, mask int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum {
				return 1 // 得到了一个合法数字
			}
			return
		}
		if !isLimit && isNum {
			p := &memo[i][mask]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}
		if !isNum { // 可以跳过当前数位
			res += f(i+1, mask, false, false)
		}
		d := 0
		if !isNum {
			d = 1 // 如果前面没有填数字，必须从 1 开始（因为不能有前导零）
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0') // 如果前面填的数字都和 n 的一样，那么这一位至多填数字 s[i]（否则就超过 n 啦）
		}
		for ; d <= up; d++ { // 枚举要填入的数字 d
			if mask>>d&1 == 0 { // d 不在 mask 中
				res += f(i+1, mask|1<<d, isLimit && d == up, true)
			}
		}
		return
	}
	return f(0, 0, true, false)
}

/***


ToDo:
digit-dp 2.0 template code

https://www.bilibili.com/video/BV1Fg4y1Q7wv/
https://leetcode.cn/problems/count-the-number-of-powerful-integers/solutions/2595149/shu-wei-dp-shang-xia-jie-mo-ban-fu-ti-da-h6ci/

时间复杂度？
 */

func digit_dp_template_2(low, high int) int {

	digitDP := func(low, high int, sumUpper int) int {
		lowS := strconv.Itoa(int(low))
		highS := strconv.Itoa(int(high))
		n := len(highS)
		lowS = strings.Repeat("0", n-len(lowS)) + lowS // 对齐

		cache := make([][]int, n)
		for i := range cache {
			cache[i] = make([]int, sumUpper+1)
			for j := range cache[i] {
				cache[i][j] = -1
			}
		}

		// 第一种写法 （前导零不影响答案）
		var f func(int, int, bool, bool) int
		f = func(p, sum int, limitLow, limitHigh bool) (res int) {
			if p == n {
				if sum > sumUpper {
					return 0
				}
				return 1
			}
			if !limitLow && !limitHigh {
				dv := &cache[p][sum]
				if *dv > 0 {
					return *dv
				}
				defer func() { *dv = res }()
			}

			lo := 0
			if limitLow {
				lo = int(lowS[p] - '0')
			}
			// 注：不要修改这里！如果对数位有其它限制， 应当写在下面的 for 循环中。
			hi := 9
			if limitHigh {
				hi = int(highS[p] - '0')
			}

			for d := lo; d <= hi; d++ {
				res += f(p+1, sum+d, limitLow && d == lo, limitHigh && d == hi)
			}

			return
		}
		//ans := f(0, 0, true, true)
		//return ans

		// 第二种写法（前导零影响答案）
		// 对于需要判断/禁止前导零的情况，可以加一个额外的维度 isNum，表示已经填入了数字（没有前导零的合法状态），最后 p=n 的时候可以根据情况返回 1 或者 0
		// 下面是计算每个数都出现偶数次的方案数
		var dfs func(int, int, bool, bool, bool) int
		dfs = func(p, mask int, limitLow, limitHigh, isNum bool) (res int) {
			if p == n {
				if !isNum {
					return 0
				}
				if mask > 0 { // 业务上的非法答案
					return 0
				}
				return 1
			}

			if !limitLow && !limitHigh {
				dv := &cache[p][mask]
				if *dv > 0 {
					return *dv
				}
				defer func() { *dv = res }()
			}

			if !isNum && lowS[p] == '0' { // 什么也不填
				res += dfs(p+1, mask, true, false, false)
			}

			lo := 0
			if limitLow {
				lo = int(lowS[p] - '0')
			}
			// 注：不要修改这里！如果对数位有其它限制，应当写在下面 for 循环中
			hi := 9
			if limitHigh {
				hi = int(highS[p] - '0')
			}
			d := lo
			if !isNum {
				d = max(lo, 1)
			}
			for ; d <= hi; d++ {
				res += dfs(p+1, mask^1<<d, limitLow && d == lo, limitHigh && d == hi, true)
			}
			return
		}

		ans := dfs(0, 0, true, true, false)
		return ans
	}

}
