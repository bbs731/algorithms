package subsequenceAutomation

func _() {

	/***
	自己写的板子。 没有使用灵神的， 灵神的板子，在处理重复的 s 或者重复的t 的时候，有问题(第一个字符的比较需要特殊处理，循环重复的
	时候处理难度太大了）。
	 */
	subsequenceAutomation := func(s string) {
		//build nxt
		// nxt[i][j] 表示在i的右侧, 字符 j (第一次出现）的最近位置
		pos := [26]int{}
		for i := range pos {
			pos[i] = len(s) + 1
		}
		nxt := make([][26]int, len(s)+1)
		for i := len(s) - 1; i >= 0; i-- {
			nxt[i+1] = pos
			pos[s[i]-'a'] = i + 1
		}
		nxt[0] = pos // 灵神的版本 是没有保存最后一个 pos的, 处理的方法，留在了 match function 去比较一下 t[0] == s[0]

		// 在 match 里
		match := func(t string) int {
			if t == "" || s == "" {
				return 0
			}
			i, j := 0, 0

			for ; j < len(t); j++ {
				i = nxt[i][t[j]-'a']
				if i == len(s)+1 {
					// 很多题，可以在这里处理，想到到的结果。
					// 如果需要重复比较 s 这里可以重置 i = 0 然后 cnts1++ 继续比较
					break
				}
			}

			//if j == len(t) {
			//	found a match for t
			//  如果需要重复比较，把 j = 0 然后 cnts2++ 继续。
			//}
			return j
		}
	}

}
