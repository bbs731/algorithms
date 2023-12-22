package 快速幂

// a^b  = power(a, b)
func binpow(a, b int) int {
	if b == 0 {
		return 1
	}
	res := binpow(a, b/2)
	if b&1 == 1 {
		return res * res * a
	}
	return res * res
}

func binpow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = res * a
		}
		a = a * a
		b >>= 1
	}
	return res
}
