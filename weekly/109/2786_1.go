package weekly

/***
视频讲解：
https://www.bilibili.com/video/BV1AM4y1x7r4/?spm_id_from=333.788&vd_source=84c3c489cf545fafdbeb3b3a6cd6a112
 */

func maxScore(nums []int, x int) int64 {
	n := len(nums)
	cache := make([][2]int, n)
	for i := range cache {
		cache[i][0] = -1
		cache[i][1] = -1
	}

	var dfs func (int, int) int
	// 这时候 j 的定义，就是 一定要选了，和前面的 dfs 的 j 的意思就变了。
	dfs = func(i int, j int) int {
		if i == n {
			return 0
		}
		if cache[i][j] != - 1{
			return cache[i][j]
		}

		res := 0
		if nums[i]%2 != j {
			// 想选， 但是不能选。
			res =dfs(i+1, j)
			cache[i][j] = res
			return res
		}

		res = max(dfs(i+1, j)+nums[i], dfs(i+1, j^1) +nums[i]-x)
		cache[i][j]= res
		return res
	}

	// 这道题的，题目要求， 0 必须选。
	return int64(dfs(0, nums[0]%2))
}

/***
按照上面的 dfs 解法， 翻译成递推。
 */

func maxScore(nums []int, x int) int64 {
	// f[i][j] = max( f[i+1][j] + nums[i], f[i+1][j^1] + nums[i] -x ) if j == nums[i] %2
	n := len(nums)
	f := [2]int{}

	for i := n-1; i>=0; i-- {
		for j:=0; j<2; j++ {
			if nums[i] %2 == j {
				f[j] = max(f[j]+nums[i], f[j^1]+nums[i]-x)
			}
		}
	}
	// 这道题，犯错，来这这里。 需要返回的是不是 f0
	return int64(f[nums[0]%2])
}

func maxScore(nums []int, x int) int64 {
	// f[i][j] = max( f[i+1][j] + nums[i], f[i+1][j^1] + nums[i] -x ) if j == nums[i] %2
	n := len(nums)
	f := [2]int{}

	for i := n-1; i>=0; i-- {
		//for j:=0; j<2; j++ {
		 	v := nums[i]
			//if nums[i] %2 == j {
			// 因为 j == v %2, 所以把所有的 j 替换成  v%2
			f[v%2] = max(f[v%2]+nums[i], f[v%2^1]+nums[i]-x)
			//}
		//}
	}
	// 这道题，犯错，来这这里。 需要返回的是不是 f0
	return int64(f[nums[0]%2])
}

