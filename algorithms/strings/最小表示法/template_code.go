package small_representation

/***
代码来自：
https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/strings.go#L343
 */
func _() {

	/***
	我自己的板子， 结合了 灵神的 + oi-wiki 上的 和 1163 题目总结的
	 */
	smallestRepresentation := func(s string) string {
		n := len(s)
		s += s
		i, k, j := 0, 0, 1
		// 循环不变量 i <= j
		for ; k < n && j+k < n; {
			if s[i+k] == s[j+k] { // 注意这里是 if
				k++
			} else {
				if s[i+k] > s[j+k] { // 改成 > 则返回字典序最大的
					// j 到 j+k 都不会是最小串的开头位置
					j += k + 1
				} else {
					// i 到 i+k 都不会是最小串的开头位置
					i, j = j, max(j, i+k)+1
				}
				k = 0 // 别忘了重置 0 因为 s[i+k] 和 s[j+k] 不相等了
			}
		}
		return s[i : i+n]
	}

	smallestRepresentation := func(s string) string {
		n := len(s)
		s = s + s // double 的快乐
		// 如果要返回一个和原串不同的字符串，初始化 i=1, j=2. 这是什么意思？ 去找下例子
		i := 0
		for j := 1; j < n; j++ {
			// 循环不变量是 i <= j
			k := 0 // 这个初始化重要，不要放在loop上避免只初始化一次。
			for k < n && s[i+k] == s[j+k] { // k < n 的条件防止越界
				k++
			}

			if k == n {
				break
			}
			if s[i+k] < s[j+k] { //改成  > 则返回字典序最大的。
				// j 到 j+k+1 都不会是最小串的开头位置
				j = j + k + 1
			} else {
				// i 到 i+k+1 都不会是最小串的开头位置
				//i, j = j, max(j, i+k)+1
				i, j = j, max(j+1, i+k+1) // j need to move forward
			}
		}
		return s[i : i+n]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
