package sliding_window


// 能转换成一个 sliding window 的问题， 太难想了。
// https://leetcode.cn/problems/replace-the-substring-for-balanced-string/solutions/2108358/tong-xiang-shuang-zhi-zhen-hua-dong-chua-z7tu/
// 灵神的想法太绝了啊。
// 枚举 右端点， windows 之外的元素的 cnts 不能 > n/4 不然无论如何，也做不到平衡，因为， windows 之外的不能修改，只修改 windows 内的元素。

// 这道题，没做过也是一个死
func balancedString(s string) int {
	n := len(s)
	m := n / 4
	ans := n + 1

	cnts := make(map[int32]int)
	for _, c := range s {
		cnts[c]++
	}
	if cnts['Q'] <= m && cnts['W'] <= m && cnts['E'] <= m && cnts['R'] <= m {
		return 0
	}
	left := 0
	// 这个窗口的逻辑还挺难写啊， 我咋就折腾了快一个小时了。
	for i, c := range s {
		cnts[c]--
		for cnts['Q'] <= m && cnts['W'] <= m && cnts['E'] <= m && cnts['R'] <= m {
			ans = min(ans, i - left+1)
			cnts[int32(s[left])]++
			left++
		}
	}
	return ans
}


	// 下面我自己的想法是错误的。
	// 终于找到自己软肋的问题了。

	//cnts := make(map[byte]int)
	//cnts['Q'] = n/4
	//cnts['W'] = n/4
	//cnts['E'] = n/4
	//cnts['R'] = n/4
	//
	//start := -1
	//for i := range s {
	//	cnts[s[i]]--
	//	if cnts[s[i]] < 0 {
	//		start = i
	//		break
	//	}
	//}
	//if  start == -1 {
	//	return 0
	//}
	//
	//
	//i :=n-1
	//for i>= 0{
	//	cnts[s[i]]--
	//	if cnts[s[i]] <0 {
	//		break
	//	}
	//	i--
	//}
	//return i - start + 1
