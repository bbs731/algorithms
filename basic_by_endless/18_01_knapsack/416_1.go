package dp

/*
let's change to dp
dfs(i,c) := dfs(i-1, c) || dfs(i-1, c-nums[i])

f[i][c] = f[i-1][c] || f[i-1][c-nums[i])

f[i+1][c] = f[i][c] || f[i][c-nums[i]]
降维：
f[c] = f[c - nums[i]]  // 根据状态方程，需要反向遍历数组。
 */

// 直接写了，降维之后的答案，没写 DP二维数组的过度代码。 赞！
func canPartition(nums []int) bool {
	n := len(nums)
	s := 0
	for _, v := range nums {
		s += v
	}
	if s%2 == 1 {
		return false
	}
	s = s / 2

	f := make([]bool, s+1)
	f[0] = true

	for i := 0; i < n; i++ {
		for j := s; j >= nums[i]; j-- {
			f[j] = f[j] || f[j-nums[i]] // 这里第一次 submit 的时候没写对 ||
		}
	}
	return f[s]
}

func canPartition_cache(nums []int) bool {
	n := len(nums)
	s := 0
	for _, v := range nums {
		s += v
	}
	if s%2 == 1 {
		return false
	}
	s = s / 2

	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, s+1)
		for j := 0; j <= s; j++ {
			cache[i][j] = -1
		}
	}

	var dfs func(int, int) bool
	dfs = func(i int, c int) bool {
		if i < 0 {
			if c == 0 {
				return true
			}
			return false
		}

		if c < 0 {
			return false
		}

		if cache[i][c] != -1 {
			return cache[i][c] == 1
		}

		ans := dfs(i-1, c) || dfs(i-1, c-nums[i])
		if ans == true {
			cache[i][c] = 1
		} else {
			cache[i][c] = 0
		}
		return ans
		//if dfs(i-1, c) {
		//	return true
		//}
		//return dfs(i-1, c-nums[i])
	}
	return dfs(n-1, s)
}

func canPartition_dfs(nums []int) bool {
	n := len(nums)
	s := 0
	for _, v := range nums {
		s += v
	}
	if s%2 == 1 {
		return false
	}
	s = s / 2

	var dfs func(int, int) bool
	dfs = func(i int, c int) bool {
		if i < 0 {
			if c == 0 {
				return true
			}
			return false
		}

		if c < 0 {
			return false
		}

		if dfs(i-1, c) {
			return true
		}
		return dfs(i-1, c-nums[i])
	}
	return dfs(n-1, s)
}
