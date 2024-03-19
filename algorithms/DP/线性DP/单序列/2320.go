package dp

func countHousePlacements(n int) int {
	MOD := int(1e9) + 7
	f := make([]int, n+10)
	f[1] = 2
	f[2] = 3
	for i:=3; i<=n; i++ {
		f[i] = f[i-1]+f[i-2]
	}

	return f[n]*f[n]%MOD
}
