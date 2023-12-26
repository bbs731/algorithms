package _600_1900

/*
1720 的分数， 还是有点难度对吗？ 很好的面试题。感觉不需要 cache 的话，应该做不了面试题。
 */
// dfs + cache
// 这道题，不用加 cache 就能过， 数据规模还是很小！
func maxLength(arr []string) int {
	n := len(arr)
	cache := make(map[string]map[int]int)
	var dfs func(string, int) int
	dfs = func(s string, i int) int {
		if i == n {
			return len(s)
		}

		if cc, ok := cache[s]; ok {
			if _, ok := cc[i]; ok {
				return cc[i]
			}
		}
		hs := make(map[byte]struct{})
		for k:=0; k<len(s); k++ {
			hs[s[k]] = struct{}{}
		}
		ha := make(map[byte]struct{})
		legal := true
		for _, c := range arr[i]{
			if _, ok := hs[byte(c)];ok {
				legal = false
			}
			if _, ok := ha[byte(c)]; ok {
				legal = false
			} else {
				ha[byte(c)]= struct{}{}
			}
		}
		ans := 0
		if legal {
			// 选  arr[i]
			//ans = len(s) + len(arr[i]) + dfs(s + arr[i], i+1)  // 这种方式， s 不是加了好几遍吗？
			ans = dfs(s + arr[i], i+1)
		}
		// 不选 arr[i]
		//ans = max(ans, len(s) + dfs(s, i+1))
		ans = max(ans, dfs(s, i+1))

		if _, ok := cache[s]; !ok {
			cache[s] = make(map[int]int)
		}
		cache[s][i]= ans
		return ans
	}
	return dfs("", 0)
}
