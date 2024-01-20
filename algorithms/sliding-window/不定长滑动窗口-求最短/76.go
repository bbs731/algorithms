package sliding_window

// 这道题，和面试17.18 非常的相似， 不过那道题是数字， 这道题是字符串。 相比之下还是更有难度， 需要比较字母出现的次数足够，所以需要一个额外高效一点的
// checker map[int32]struct{} 如果是纯粹的 map contains 关系， 恐怕会时间上超时。

// 当命运的车轮滚滚向前，谁都逃不过这时代给与的冲击。

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
				if tmp[lc] < tm[lc] {  // 这里比较 tm[lc] 不是 tm[c] 第一次提交WA在这里。
					delete(checker, lc)
				}
			}
			left++
		}
	}

	return ans
}
