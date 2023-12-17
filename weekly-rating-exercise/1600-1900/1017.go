package _600_1900


func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	if n ==1 {
		return "1"
	}

	res := make([]byte, 0)

	/*  左右进制的原理， 都是一样的。
	https://leetcode.cn/problems/convert-to-base-2/solutions/2209807/fu-er-jin-zhi-zhuan-huan-by-leetcode-sol-9qlh/

	看官方的题解!`
	 */
	for n !=0 {
		remainder := n &1
		res = append(res, byte('0' + remainder))
		n -= remainder
		n /=-2
	}

	l := len(res)
	for i:=0; i<= len(res)/2; i++ {
		res[i], res[l-1-i] = res[l-1-i], res[i]
	}
	return string(res)
}
