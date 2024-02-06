package subsequenceAutomation

func _() {

	subsequenceAutomation := func(s string) {
		//build nxt
		// nxt[i][j] 表示在i的右侧, 字符 j (第一次出现）的最近位置
		pos := [26]int{}
		for i := range pos {
			pos[i] = len(s) // 初始化， mark end, 在 match 的时候会判断使用
		}
		nxt := make([][26]int, len(s))

		for i := len(s) - 1; i >= 0; i-- {
			nxt[i] = pos
			pos[s[i]-'a'] = i // 其实是为了 nxt[i-1] 准备的 pos 值。
			// 这样写边界处理的很清晰， 但是有一个问题就是 nxt[0] = pos[s[1]-'a'], 我们发现 没有 nxt[-1] 也就是说
			// pos[s[0]-'a'] 的这个 match 关系我们没有保存，是丢掉了的。针对这个问题，就需要在 match 里对 s[0] 特殊处理。
		}

		// 在 match 里
		match := func(t string) int {
			if t == "" || s == "" {
				return 0
			}
			i, j := 0, 0
			// 因为我们在 build nxt 的时候，丢弃了 pos[s[0]-'a'] 这个信息，所以需要在这里，特殊处理一下。
			if t[0] == s[0] {
				j = 1 // t[0] 匹配。
			}

			for ; j < len(t); j++ {
				i = nxt[i][t[j]-'a']
				if i == len(s) {
					// 很多题，可以在这里处理，想到到的结果。
					break
				}
			}

			//if j == len(t) {
			//	found a match
			//}
			return j
		}
	}

}
