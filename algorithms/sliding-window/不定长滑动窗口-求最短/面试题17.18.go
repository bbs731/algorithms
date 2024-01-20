package sliding_window

// 为啥不能一次秒杀？ WA 了两次， 一次逻辑检查， 一次边界条

func shortestSeq(big []int, small []int) []int {
	sm := make(map[int]struct{})
	tmp := make(map[int]int)
	for _, s := range small {
		sm[s]= struct{}{}
	}

	ans := len(big) + 1  // 这里没 + 1 造成 [1,2,3]  [1,2,3] 的测试用例过不去。
	result:= []int{}


	left := 0
	for i, b := range big {
		// not interested number
		if _, ok := sm[b]; !ok {
			continue
		}

		tmp[b]++
		for len(tmp) == len(sm) {
			// save answer if possible
			if i - left + 1 < ans {
				ans = i - left + 1
				result = []int{left, i}
			}

			// 这里更新 tmp 的时候，忘记检查是否是在 small map 里面了。
			if _, ok := sm[big[left]]; ok {
				l := big[left]
				tmp[l]--
				if tmp[l] == 0 {
					delete(tmp,l)
				}
			}
			left++
		}
	}
	return result
}
