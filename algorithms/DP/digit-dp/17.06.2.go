package digit_dp

import "strconv"

/***
https://leetcode.cn/problems/number-of-2s-in-range-lcci/solutions/1750395/by-endlesscheng-x4mf/

灵神给的模版代码， 可理解， 可重用。 oi-wiki 上的换一道题，不知道怎么弄。
*/


func numberOf2sInRange(n int) int {

	// step 1
	// convert to string 如果是 10进制的题目
	// m := len(s)
	// 计算二进制的长度， 如果是二进制题目。
	//m := bits.Len(uint(n))
	//memo := make([][]int, m)

	s := strconv.Itoa(n)
	m := len(s)
	memo := make([][]int, m)
	for i := range memo {
		memo[i]= make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}

	var f func(int, int, int, bool, bool) int
	// mask 不是一定的， 有的题目不需要 mask, 需要 cnt1 譬如 count 1 出现的次数。
	f = func(i, mask, cnt1 int, isLimit, isNum bool) (res int) {
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
			res += f(i+1, mask, cnt1, false, false)
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
			//if mask>>d&1 == 0 { // d 不在 mask 中
			c := cnt1

			if d== 2 {  // count digit 2,  和 count digit 1 唯一的区别，就在这里。
				c++
			}
			res += f(i+1, mask|1<<d, c, isLimit && d == up, true)
			//}
		}
		return
	}
	return f(0, 0,0, true, false)
}


