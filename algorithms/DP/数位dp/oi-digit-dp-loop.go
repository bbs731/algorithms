package leetcode

/*
https://oi-wiki.org/dp/number/

LC233
https://leetcode.cn/problems/number-of-digit-one/description/

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
	// 把 n break into digits array, index start from 1
	l := 0
	for n > 0 {
		l++
		a[l] = n % 10
		n = n / 10
	}

	// 作为分析，我们考虑  n = 543, i= 3 时 a[i] = 5 的情况
	for i := l; i >= 1; i-- {
		// 考虑的是 1 到 i-1 位的情况， 都是满位的情况。
		for j := 0; j < 10; j++ {
			result[j] += a[i] * dp[i-1] // 这里考虑的是后 i-1 满位，对于 0~9 数字的贡献， 对每个数字贡献值一样的，  i= 3 时 a[i]= 5,  考虑的是 a[3] = 0 or 1 or 2 or 3 or 4 对于后 i-1 为的贡献
		}

		// 考虑的是 第 i 位的情况， 但是不包括最高位的那个值 a[3]=5
		for j := 0; j < a[i]; j++ {
			result[j] += mi[i-1] // 这里考虑的是 第 i 位， 当 i= 3, digit 是 5 时，对于数字 0 ~ 4 的贡献， 其实这里，可以只考虑 1 ~ 4 就不需要后面处理前导0 了。 譬如 3xx,  那么对数字3 的贡献就是  mi[2] =100
		}

		// 这里考虑的是最高位的那个 digit 5 就一个值 a[3]=5
		tmp -= mi[i-1] * a[i]   // 这里考虑的是，贴着上界的情况，就是 a[3] = 5 的情况， 对于 5 的数字贡献就是  43 +1 因为  tmp=543  tmp - a[3]* 100 = 43
		result[a[i]] += tmp + 1 // tmp + 1 的意思，按照例子就是  0 到 43 一共 44次。

		/***
			https://oi-wiki.org/dp/number/:
			最后考虑下前导零，第 i 位为前导 0 时，此时 1 到 i-1 位也都是 0，也就是多算了将 i-1 位填满的答案，需要额外减去。

			上面这句话，理解了无数次，我还是觉得说的有误  " 1 到 i-1 位也都是 0" 这个点很荒谬啊。
			我的理解是: 第 ith 位 如果是前导 0 的话， 我们多贡献了， 10^(i-1) 的 数字 0 ，需要从 result[0] 中去除。

		 */
		result[0] -= mi[i-1] //这里，如果考虑 第i位的时候， 不考虑 0， 这里其实可以不处理这个前导 0 (看下面的版本）
	}

	return result[1]
}

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
	tmp := n
	l := 0
	for n > 0 {
		l++
		a[l] = n % 10
		n = n / 10
	}

	for i := l; i >= 1; i-- {
		for j := 0; j < 10; j++ {
			result[j] += a[i] * dp[i-1]
		}

		for j := 1; j < a[i]; j++ { // 这里 j 从 1 开始，不统计0， 下面不用处理前导0 是和上面版本的区别。
			result[j] += mi[i-1]
		}
		tmp -= mi[i-1] * a[i]
		result[a[i]] += tmp + 1
	}
	return result[1]
}
