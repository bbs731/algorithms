package 快速幂

// a^b
func binpow(a, b int) int {
	if b == 0 {
		return 1
	}
	res := binpow(a, b/2)
	if b&1 == 1 {
		return res * res * b
	}
	return res * res
}
