package digit_dp

import (
	"math/bits"
	"strconv"
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

func digit_dp_template (n int) int {

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