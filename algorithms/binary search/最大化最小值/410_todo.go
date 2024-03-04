package binary_search

/***
可以做答案的隐式的二分， 原因， 是这样的， 对于给定k个分组。 如果，所求的 ans 越小，越不容易成功， 如果 ans 越大，越容易成功，
ans 的分布有单调性， 所以可以在 ans 上面进行二分答案。
 */

func splitArray(nums []int, k int) int {
	check := func(mid int) bool {
		cnt := 1
		s := 0
		for _, x := range nums {
			if s+x <= mid {
				s += x
			} else {
				cnt++
				if cnt > k {
					return false
				}
				s = x
			}
		}
		return true
	}

	//n := len(nums)
	mn, s := 0, 0
	for _, x := range nums {
		mn = max(mn, x)
		s += x
	}
	l, r := mn-1, s

	for l+1 < r {
		mid := (l + r) >> 1

		if check(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

/***
let's do it in dp

f[i][k] = min(f[i][k], max(f[j][k-1],  sum(j+1, i))) for j = [k-1... i]
 */

func splitArray(nums []int, k int) int {
	n := len(nums)
	f := make([][]int, n+1)
	inf := int(1e15)
	for i := range f {
		f[i] = make([]int, k+1)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	// 计算一下，前缀和
	ps := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ps[i] = ps[i-1] + nums[i-1]
	}

	// 难点就在初始化， 需要让 i -> i+1 多留一个 f[0][0] 的状态，初始化成  f[0][0] = 0

	// 你为了初始化成 f[0][0] 所以 f[i] i 的 index 需要从 1 开始， 给 0 留出来位置。
	f[0][0] = 0
	// 初始化 f
	//for i := 0; i < n; i++ {
	//	f[i][1] = ps[i+1]
	//}

	for i := 0; i < n; i++ {
		for j := 1; j <= k; j++ {
			for p := j - 1; p <= i; p++ {
				f[i+1][j] = min(f[i+1][j], max(f[p][j-1], ps[i+1]-ps[p]))
			}
		}
	}

	return f[n][k]
}
