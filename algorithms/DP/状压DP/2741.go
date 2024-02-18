package dp

func specialPerm(nums []int) int {
	n := len(nums)
	const mod int = 1e9 + 7
	m := 1 << n

	//cache := make([][1 << 15]int, n)
	cache := make([][1 << 15]int, n)
	for i := range cache {
		for j := 0; j < 1<<15; j++ {
			cache[i][j] = -1
		}
	}

	var dfs func(int, int) int
	// i stands for the last element added. 如果是这样的话， perm 就不需要了。
	//dfs = func(i, bitmask int, perm []int) int {
	dfs = func(i, bitmask int) int {
		if bitmask == m-1 {
			return 1
		}

		if cache[i][bitmask] != -1 {
			return cache[i][bitmask]
		}

		res := 0
		// 如果 i 代表 last digit added, 那么就不需要 perm 了。
		//last := perm[len(perm)-1]
		last := nums[i]

		for k := 0; k < n; k++ {
			if bitmask&(1<<uint(k)) != 0 {
				continue
			}
			if last%nums[k] == 0 || nums[k]%last == 0 {
				//perm = append(perm, nums[k])
				res += dfs(k, bitmask|1<<uint(k))
				res %= mod
				//restore
				//perm = perm[:len(perm)-1]
			}

		}

		cache[i][bitmask] = res % mod
		return res % mod
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += dfs(i, 1<<uint(i))
		ans %= mod
	}
	return ans
}
