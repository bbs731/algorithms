package permuation

/***
一次就能写对， 赞啊！
 */
func permute(nums []int) (ans [][]int) {
	n := len(nums)

	var dfs func(int, int, []int)
	dfs = func(i int, mask int, l []int) {
		if i == n {
			cp := make([]int, n)
			copy(cp, l)
			ans = append(ans, cp)
			return
		}
		for j := 0; j < n; j++ {
			if mask&(1<<j) != 0 {
				continue
			}
			l[i] = nums[j]
			dfs(i+1, mask|1<<j, l)
			// undo l
		}
	}

	dfs(0, 0, make([]int, n))
	return
}
