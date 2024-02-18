package weekly

/***
给你一个下标从 0 开始的整数数组 nums ，它包含 n 个 互不相同 的正整数。如果 nums 的一个排列满足以下条件，我们称它是一个特别的排列：

对于 0 <= i < n - 1 的下标 i ，要么 nums[i] % nums[i+1] == 0 ，要么 nums[i+1] % nums[i] == 0 。
请你返回特别排列的总数目，由于答案可能很大，请将它对 109 + 7 取余 后返回。



示例 1：

输入：nums = [2,3,6]
输出：2
解释：[3,6,2] 和 [2,6,3] 是 nums 两个特别的排列。
示例 2：

输入：nums = [1,4,3]
输出：2
解释：[3,1,4] 和 [4,1,3] 是 nums 两个特别的排列。


提示：

2 <= nums.length <= 14
1 <= nums[i] <= 10^9

 */

func specialPerm(nums []int) int {
	n := len(nums)
	const mod int = 1e9 + 7
	m := 1 << uint(n)

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

/***
下面的这个例子是错的。 因为 （i，bitmask) 的组合定义状态，不唯一。 为什么不唯一呢？ 举一个例子： ????

譬如，该填 第3个位置了， i = 3   bitmask 00011 这个状态，可能是  000,n[1],n[0]  也可能是 000, n[0],n[1] 不能唯一表示。 因为 bitmask 只表示了取还是没取， 但是没有办法确定位置。 如果 i 也确定不了唯一的位置信息，那么这个状态，就不是唯一的。 (赞啊！问题逻辑分析的很清楚啊！）


这个，为什么错误了， 加了 cache 就错了。 搞不懂？ cache 的这种形式为什么不行？ 应该就是状态，不唯一的问题。下面定义 i the bitmask 的状态不能唯一表示一个状态！
 */
func specialPerm(nums []int) int {
	n := len(nums)
	//const mod int = 1e9 + 7

	//cache := make([][1 << 15]int, n)
	cache := make([][]int, n)
	m := 1 << n
	for i := range cache {
		cache[i] = make([]int, m)
		for j := 0; j < m; j++ {
			cache[i][j] = -1
		}
	}
	var dfs func(int, int, []int) int
	dfs = func(i, bitmask int, perm []int) int {
		if i == n {
			return 1
		}

		if cache[i][bitmask] != -1 {
			return cache[i][bitmask]
		}

		res := 0
		last := perm[len(perm)-1]

		for k := 0; k < n; k++ {
			if bitmask&(1<<uint(k)) != 0 {
				continue
			}
			if last%nums[k] == 0 || nums[k]%last == 0 {
				perm = append(perm, nums[k])
				res += dfs(i+1, bitmask|1<<uint(k), perm)
				//restore
				perm = perm[:len(perm)-1]
			}

		}

		cache[i][bitmask] = res
		return res
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += dfs(1, 1<<uint(i), []int{nums[i]})
	}
	return ans

}
