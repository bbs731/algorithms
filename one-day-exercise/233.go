package one_day_exercise

import "strconv"

/***

用一个 digit dp 最通用的模版， 实验一下。
虽然有点杀鸡用牛刀了。 但是为了学习一下通用模版。

***/


func countDigitOne(n int) int {
	s := strconv.Itoa(n)
	m := len(s)

	memo := make([][]int, m)
	for i := range memo {
		memo[i]= make([]int,m)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var f func(int, int, bool, bool) int
	// mask 不一定需要
	f = func(i,  cnt1 int, isLimit, isNum bool) (res int) {
		if i == m {
			//if isNum {
			//	return 1 // 得到了一个合法数字
			//}
			return cnt1
		}
		if !isLimit && isNum {
			p := &memo[i][cnt1]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}
		if !isNum { // 可以跳过当前数位
			res += f(i+1,cnt1, false, false)
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
			c := cnt1
			if d == 1 {
				c++
			}
			res += f(i+1, c, isLimit && d == up, true)
		}
		return
	}
	return f(0, 0, true, false)
}

