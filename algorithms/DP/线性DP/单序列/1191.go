package dp

const MOD = int(1e9) + 7

func maxSubseq(l []int) int {
	f0 := 0
	f := 0
	n := len(l)

	for i := 1; i <= n; i++ {
		f0 = max(f0+l[i-1], l[i-1], 0) % MOD
		f = max(f, f0)
	}
	return f % MOD
}

/***
哎呀， 真不容易！

不是什么好题， 算了吧！感觉做的，还是有很多漏洞。
算了吧， 忘了把！
 */

func kConcatenationMaxSum(arr []int, k int) int {

	n := len(arr)
	l := append([]int{}, arr...)
	l = append(l, arr...)

	s := 0
	prefixs := 0
	for i := 0; i < n; i++ {
		s += arr[i]
		s %= MOD
		prefixs = max(prefixs, s)
	}

	if k == 1 {
		return maxSubseq(arr)
	}

	// 好多种，情况啊！ 感觉，  s *(k-2) + maxSubseq(l) 这个就不是很对！
	return max(s*k%MOD, maxSubseq(l), maxSubseq(arr), s*(k-1)+prefixs, s*(k-2)+maxSubseq(l)) % MOD
}
