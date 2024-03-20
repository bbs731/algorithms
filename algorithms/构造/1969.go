package constructive

/****
终于看到了，构造题，是长什么样子的！
 */
const MOD = int(1e9) + 7

func minNonZeroProduct(p int) int {
	p2 := 1 << p

	return (p2 - 1) % MOD * power(p2-2, 1<<(p-1)-1) % MOD
	// 这个不行， power 是 MOD之后的结果，所以计算 2^(p-1) 最好用 bit shift
	//return (p2 - 1) % MOD * power(p2-2, power(2, p-1)-1) % MOD
}

func power(x, n int) int {
	ans := 1
	for x %= MOD; n > 0; n = n >> 1 {
		if n&1 > 0 {
			ans = ans * x % MOD
		}
		x = x * x % MOD
	}

	return ans
}
