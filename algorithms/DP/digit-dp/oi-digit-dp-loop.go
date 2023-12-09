package leetcode

/*
https://oi-wiki.org/dp/number/

example 1,  方法1
 */

const N = 15

/* 思考的例子，如 n = 543 */
func countDigitOne(n int) (ans int) {

	mi := make([]int, N)
	dp := make([]int, N)
	a := make([]int, N)
	result := make([]int, 10)
	mi[0] = 1

	for i := 1; i < N; i++ {
		dp[i] = 10*dp[i-1] + mi[i-1]
		mi[i] = 10 * mi[i-1]
	}

	// solve
	tmp := n
	l := 0
	for n > 0 {
		l++
		a[l] = n % 10
		n = n / 10
	}

	// 作为分析，我们考虑  n = 543, i= 3 时的情况
	for i := l; i >= 1; i-- {
		for j := 0; j < 10; j++ {
			result[j] += a[i] * dp[i-1] // 这里考虑的是后 i-1 满位，对于 0~9 数字的贡献， 对每个数字贡献值一样的，  i= 3 时 a[i]= 5,  考虑的是 a[3] = 0 or 1 or 2 or 3 or 4 对于后 i-1 为的贡献
		}
		for j := 0; j < a[i]; j++ {
			result[j] += mi[i-1] // 这里考虑的是 第 i 位， 当 i= 3 时，对于数字 0 ~ 4 的贡献， 其实这里，可以只考虑 1 ~ 4 就不需要后面处理前导0 了。 譬如 3xx,  那么对数字3 的贡献就是  mi[2] =100
		}
		tmp -= mi[i-1] * a[i] // 这里考虑的是，贴着上界的情况，就是 a[3] = 5 的情况， 对于 5 的数字贡献就是  43 +1 因为  tmp=543  tmp - a[3]* 100 = 43
		result[a[i]] += tmp + 1
		result[0] -= mi[i-1] //这里，如果考虑 第i位的时候， 不考虑 0， 这里其实可以不处理这个前导 0
	}

	return result[1]
}
