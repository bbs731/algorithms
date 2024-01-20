package sliding_window

/***

LCR 017. 最小覆盖子串
https://leetcode.cn/problems/M1oyTv/description/
 */


func minWindow(s string, t string) string {
	tm := make(map[int32]int)
	for _, tc := range t {
		tm[tc]++
	}
	ans := ""
	left := 0
	tmp := make(map[int32]int)
	checker := make(map[int32]struct{})
	for i, c := range s {
		if _, ok := tm[c]; !ok {
			continue
		}
		tmp[c]++
		// 只有当数目够了，才把 c 加入到 checker
		if tmp[c] >= tm[c] {
			checker[c] = struct{}{}
		}

		for len(checker) == len(tm) {
			// save answer
			if ans == "" {
				ans = s[left:i+1]
			} else if len(ans) > i-left+1 {
				ans = s[left:i+1]
			}
			// 移除 left element
			lc := int32(s[left])
			if _, ok := tm[lc]; ok {
				tmp[lc]--
				if tmp[lc] < tm[lc] {
					delete(checker, lc)
				}
			}
			left++
		}
	}

	return ans
}
